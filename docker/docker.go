package docker

import (
	"github.com/docker/docker/client"

	"context"
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
	cli, err := getClient()
	if err != nil {
		return err
	}
	return cli.ContainerKill(context.Background(), containerID, "")
}
