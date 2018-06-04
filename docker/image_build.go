package docker

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/muka/mufaas/asset"
	"github.com/muka/mufaas/template"
	"github.com/muka/mufaas/util"
	log "github.com/sirupsen/logrus"
)

// ImageBuildOptions options to build an image
type ImageBuildOptions struct {
	Name       string
	Type       string
	TypesPath  []string
	Archive    string
	Dockerfile string
	Labels     map[string]string
}

// read Dockerfile, inject idle command and rebuild to a tar archive
func parseDockerfile(dockerFile, dst string, opts *ImageBuildOptions) (string, error) {

	log.Debugf("Read %s", dockerFile)
	b, err := ioutil.ReadFile(dockerFile)
	if err != nil {
		return "", err
	}

	// get idle bin from assets
	log.Debugf("Current GOARCH %s", runtime.GOARCH)
	assetName := "idle/bin/idle"
	switch runtime.GOARCH {
	case "arm":
	case "arm64":
	case "amd64":
		assetName += "-" + runtime.GOARCH
		break
	default:
		return "", fmt.Errorf("architecture `%s` not supported, idle binaries may need to be regenerated", runtime.GOARCH)
	}

	idleDest := filepath.Join(dst, "idle")
	assetData, err := asset.Asset(assetName)
	if err != nil {
		return "", err
	}

	if _, err1 := os.Stat(idleDest); err1 == nil {
		fmt.Print("Drop idle")
		err2 := os.Remove(idleDest)
		if err2 != nil {
			return "", err2
		}
	}

	//Create idle
	err = ioutil.WriteFile(idleDest, assetData, 0755)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	r := bufio.NewReader(bytes.NewReader(b))
	line, err := r.ReadString('\n')
	for err == nil {
		if len(line) > 3 {
			if strings.ToUpper(line[:3]) == "CMD" {
				srcCmd := base64.RawStdEncoding.EncodeToString([]byte(line[4:]))
				cmds := fmt.Sprintf("\nENV %s '%s'\nADD ./idle /idle\nCMD [\"/idle\"]\n", CmdEnvKey, srcCmd)
				_, err1 := w.WriteString(cmds)
				if err1 != nil {
					return "", err1
				}
			} else {
				_, err1 := w.WriteString(line)
				if err1 != nil {
					return "", err1
				}
			}
		}
		line, err = r.ReadString('\n')
	}
	if err != io.EOF {
		if err != nil {
			return "", err
		}
	}

	err = w.Flush()
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(dockerFile, buf.Bytes(), 0644)
	if err != nil {
		return "", err
	}

	err = util.CreateTar(dst)
	if err != nil {
		return "", err
	}

	log.Debugf("New archive created at %s.tar", dst)

	return dst + ".tar", nil
}

func extractArchive(archivePath string) (dst string, err error) {
	dst, err = ioutil.TempDir("", "mufaas_build")
	if err != nil {
		return dst, err
	}
	log.Debugf("Extract archive %s", dst)
	err = util.ExtractTar(archivePath, dst)
	if err != nil {
		return dst, err
	}
	return dst, err
}

// ImageBuild builds a docker image from the image directory
func ImageBuild(opts ImageBuildOptions) (*types.ImageSummary, error) {

	name := opts.Name
	if name == "" {
		return nil, errors.New("Image name not provided")
	}

	archive := opts.Archive
	if _, err := os.Stat(archive); os.IsNotExist(err) {
		return nil, err
	}

	if opts.Dockerfile == "" {
		opts.Dockerfile = "Dockerfile"
	}

	if opts.Labels == nil {
		opts.Labels = make(map[string]string)
	}
	opts.Labels[DefaultLabel] = "1"

	//extract
	dst, err := extractArchive(opts.Archive)
	//copy dst1 in case is overwritten later
	dst1 := dst
	defer os.Remove(dst1)
	if err != nil {
		return nil, err
	}

	//copy Dockerfile from TypesPath if not provided by the archive
	dockerFile := filepath.Join(dst, opts.Dockerfile)
	if _, err1 := os.Stat(dockerFile); os.IsNotExist(err1) {
		dst, err1 = template.CreateFunction(dst, opts.Type, opts.TypesPath)
		if err1 != nil {
			return nil, err1
		}
		dockerFile = filepath.Join(dst, "Dockerfile")
		// ensure the built dst is also removed
		if dst != dst1 {
			dst2 := dst
			defer os.Remove(dst2)
		}
	}

	// ensure we have a Dockerfile
	if _, err1 := os.Stat(dockerFile); os.IsNotExist(err1) {
		return nil, errors.New("Dockerfile not found")
	}

	archive, err = parseDockerfile(dockerFile, dst, &opts)
	if err != nil {
		return nil, err
	}
	defer os.Remove(archive)

	cli, err := getClient()
	if err != nil {
		return nil, err
	}

	log.Debugf("Build image %s from %s", name, archive)

	dockerBuildContext, buildContextErr := os.Open(archive)
	if buildContextErr != nil {
		return nil, buildContextErr
	}
	defer dockerBuildContext.Close()

	buildOptions := types.ImageBuildOptions{
		Dockerfile: opts.Dockerfile,
		// NoCache:     true,
		// PullParent:  false,
		ForceRemove: true,
		Remove:      true,
		Tags:        []string{name},
		Labels:      opts.Labels,
	}
	buildResponse, buildErr := cli.ImageBuild(context.Background(), dockerBuildContext, buildOptions)
	if buildErr != nil {
		return nil, buildErr
	}
	defer buildResponse.Body.Close()

	scanner := bufio.NewScanner(buildResponse.Body)
	var info struct {
		Stream string `json:"stream"`
	}
	for scanner.Scan() {
		if log.GetLevel() == log.DebugLevel {
			err1 := json.Unmarshal(scanner.Bytes(), &info)
			if err1 != nil {
				log.Warnf("Unmarshal failed: %s", err1.Error())
				continue
			}
			s := strings.Replace(info.Stream, "\n", "", -1)
			if len(s) == 0 {
				continue
			}
			log.Debugf(s)
		}
	}

	var imageInfo *types.ImageSummary
	filter := []string{"reference=" + name}
	imageList, err := ImageList(filter)
	if err != nil {
		return nil, err
	}
	if len(imageList) != 1 {
		return nil, fmt.Errorf("Image %s not found, build failed", name)
	}

	imageInfo = &imageList[0]
	log.Debugf("Built image %s [%s]", name, imageInfo.ID)
	return imageInfo, nil
}
