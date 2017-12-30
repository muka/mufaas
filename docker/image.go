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
	"github.com/docker/docker/api/types/filters"
	"github.com/muka/mufaas/asset"
	"github.com/muka/mufaas/util"
	log "github.com/sirupsen/logrus"
)

const CmdEnvKey = "MUFAASSOURCECMD"

type ImageBuildOptions struct {
	Name       string
	Archive    string
	Dockerfile string
	Labels     map[string]string
}

// List return built images, filtered by a list of docker
// compatible filters (key=value) eg. [id=..., name=...]
func ImageList(listFilters []string) ([]types.ImageSummary, error) {
	cli, err := getClient()
	if err != nil {
		return nil, err
	}
	f, err := buildFilter(listFilters)
	if err != nil {
		return nil, err
	}
	if log.GetLevel() == log.DebugLevel {
		strfilter, err := filters.ToParam(f)
		if err != nil {
			log.Warnf("Failed to convert params:, %s", err.Error())
		} else {
			log.Debugf("Image filters: %s", strings.Replace(strfilter, "\\\"", "\"", -1))
		}
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{
		All:     true,
		Filters: f,
	})

	if err != nil {
		return nil, err
	}
	return images, nil
}

func parseDockerfile(opts ImageBuildOptions) (string, error) {

	//extract
	dst, err := ioutil.TempDir("", "mufaas_build")
	defer os.Remove(dst)
	if err != nil {
		return "", err
	}

	log.Debugf("Extract archive %s", dst)
	err = util.ExtractTar(opts.Archive, dst)
	if err != nil {
		return "", err
	}

	// read Dockerfile
	dockerFile := filepath.Join(dst, opts.Dockerfile)
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
		return "", fmt.Errorf("Architecture `%s` not supported, idle binaries may need to be regenerated.", runtime.GOARCH)
	}

	idleDest := filepath.Join(dst, "idle")
	assetData, err := asset.Asset(assetName)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(idleDest); err == nil {
		fmt.Print("Drop idle")
		err := os.Remove(idleDest)
		if err != nil {
			return "", err
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
				srcCmd := base64.StdEncoding.EncodeToString([]byte(line[4:]))
				cmds := fmt.Sprintf("\nENV %s '%s'\nADD ./idle /idle\nCMD [\"/idle\"]\n", CmdEnvKey, srcCmd)
				_, err := w.WriteString(cmds)
				if err != nil {
					return "", err
				}
			} else {
				_, err := w.WriteString(line)
				if err != nil {
					return "", err
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

	log.Printf("%s", buf.String())

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

	archive, err := parseDockerfile(opts)
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
			err := json.Unmarshal(scanner.Bytes(), &info)
			if err != nil {
				log.Warnf("Unmarshal failed: %s", err.Error())
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

//ImageRemove remove an image
func ImageRemove(id string, force bool) (err error) {
	cli, err := getClient()
	if err != nil {
		return err
	}
	log.Debugf("Remove image %s", id)
	_, err = cli.ImageRemove(context.Background(), id, types.ImageRemoveOptions{
		Force:         force,
		PruneChildren: force,
	})
	return err
}
