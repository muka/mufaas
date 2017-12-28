package docker

import (
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

const DefaultLabel = "mufaas"

var dockerClient *client.Client

func init() {
	go func() {
		err := watchEvents()
		if err != nil {
			panic(err)
		}
	}()
}

func getClient() (*client.Client, error) {
	if dockerClient != nil {
		return dockerClient, nil
	}
	var err error
	dockerClient, err = client.NewEnvClient()
	return dockerClient, err
}

func buildFilter(listFilters []string) (filters.Args, error) {
	f := filters.NewArgs()
	f.Add("label", DefaultLabel+"=1")
	for _, filter := range listFilters {
		filterParts := strings.Split(filter, "=")
		if len(filterParts) < 2 {
			return f, fmt.Errorf("Filter `%s` should have format key=value")
		}
		f.Add(filterParts[0], strings.Join(filterParts[1:], "="))
	}
	return f, nil
}
