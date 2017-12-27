package docker

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	log "github.com/sirupsen/logrus"
)

// List return built images, filtered by a list of docker
// compatible filters (key=value) eg. [id=..., name=...]
func ImageList(listFilters []string) ([]types.ImageSummary, error) {
	cli, err := getClient()
	if err != nil {
		return nil, err
	}

	f := filters.NewArgs()
	f.Add("label", DefaultLabel+"=1")
	for _, filter := range listFilters {
		filterParts := strings.Split(filter, "=")
		if len(filterParts) < 2 {
			return nil, fmt.Errorf("Filter `%s` should have format key=value")
		}
		f.Add(filterParts[0], strings.Join(filterParts[1:], "="))
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

// ImageBuild builds a docker image from the image directory
func ImageBuild(name string, archive string) (*types.ImageSummary, error) {
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
		Dockerfile: "Dockerfile",
		// NoCache:     true,
		// PullParent:  false,
		ForceRemove: true,
		Remove:      true,
		Tags:        []string{name},
		Labels:      map[string]string{DefaultLabel: "1"},
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
