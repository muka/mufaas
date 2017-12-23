package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"context"
	"io"
	"os"
	"time"
)

const DefaultLabel = "mufaas"

type dockerInfo struct {
	Stream string `json:"stream"`
}

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

// Pull image from registry
func Pull(name string, verbose bool) error {
	cli, err := getClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	r, pullErr := cli.ImagePull(ctx, name, types.ImagePullOptions{})
	if pullErr != nil {
		return pullErr
	}

	if verbose {
		io.Copy(os.Stdout, r)
	}

	return nil
}

// Stop interrupts a running container
func Stop(containerID string) (err error) {
	cli, err := getClient()
	if err != nil {
		return err
	}
	timeout := time.Duration(1) * time.Second
	err = cli.ContainerStop(context.Background(), containerID, &timeout)
	return err
}

// Remove interrupts and remove a running container
func Remove(containerID string) (err error) {
	cli, err := getClient()
	if err != nil {
		return err
	}
	return cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: true})
}

// Kill a running container
func Kill(containerID string) (err error) {
	cli, err := getClient()
	if err != nil {
		return err
	}
	return cli.ContainerKill(context.Background(), containerID, "")
}
