package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
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

// Kill a running container
func Kill(containerID string) (err error) {
	log.Debugf("Kill container %s", containerID)
	cli, err := getClient()
	if err != nil {
		return err
	}
	return cli.ContainerKill(context.Background(), containerID, "")
}

// Remove a container
func Remove(containerID string) (err error) {
	log.Debugf("Remove container %s", containerID)
	cli, err := getClient()
	if err != nil {
		return err
	}
	return cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: true})
}
