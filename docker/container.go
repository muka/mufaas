package docker

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	log "github.com/sirupsen/logrus"
)

type ContainerCreated struct {
	ID   string
	Name string
}

type CreateOptions struct {
	Name       string
	Image      string
	Cmd        []string
	Env        []string
	Privileged bool
}

// Create a new container
func Create(opts *CreateOptions) (*ContainerCreated, error) {

	cli, err := getClient()
	if err != nil {
		return nil, err
	}

	imageName := opts.Image
	if len(imageName) == 0 {
		return nil, errors.New("Missing image name")
	}

	baseContainerName := opts.Name
	if len(baseContainerName) == 0 {
		baseContainerName = imageName
	}

	containerName := DefaultLabel + "-" + baseContainerName

	log.Debugf("Creating container %s (from %s)\n", containerName, imageName)
	containerConfig := &container.Config{
		Cmd:          opts.Cmd,
		Env:          opts.Env,
		Image:        imageName,
		AttachStdin:  false,
		AttachStderr: true,
		AttachStdout: true,
		Tty:          true,
		StdinOnce:    true,
		Labels: map[string]string{
			DefaultLabel: "1",
		},
	}

	hostConfig := &container.HostConfig{
		AutoRemove: false,
		Privileged: opts.Privileged,
	}

	netConfig := &network.NetworkingConfig{}

	ctx := context.Background()
	resp, cerr := cli.ContainerCreate(ctx, containerConfig, hostConfig, netConfig, containerName)
	if cerr != nil {
		return nil, cerr
	}

	log.Debugf("Created container %s (%s)", containerName, resp.ID)
	for _, m := range resp.Warnings {
		log.Debugf("[%s]: %s", resp.ID, m)
	}

	return &ContainerCreated{
		ID:   resp.ID,
		Name: containerName,
	}, nil
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
func Remove(containerID string, forceRemove bool) (err error) {
	log.Debugf("Remove container %s", containerID)
	cli, err := getClient()
	if err != nil {
		return err
	}
	return cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: forceRemove})
}

// List containers
func List(listFilters []string) (list []types.Container, err error) {
	cli, err := getClient()
	if err != nil {
		return nil, err
	}
	f, err := buildFilter(listFilters)
	if err != nil {
		return nil, err
	}
	options := types.ContainerListOptions{
		All:     true,
		Filters: f,
	}
	ctx := context.Background()
	list, err = cli.ContainerList(ctx, options)
	if err != nil {
		return nil, err
	}
	return list, nil
}
