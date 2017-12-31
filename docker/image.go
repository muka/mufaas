package docker

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	log "github.com/sirupsen/logrus"
)

const CmdEnvKey = "MUFAASSOURCECMD"

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
