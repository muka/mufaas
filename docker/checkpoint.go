package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	log "github.com/sirupsen/logrus"
)

// CheckpointCreate create a container checkpoint
func CheckpointCreate(containerID string) (err error) {
	log.Debugf("Create checkpoint for container %s", containerID)
	cli, err := getClient()
	if err != nil {
		return err
	}
	opts := types.CheckpointCreateOptions{
		// CheckpointDir: "",
		CheckpointID: containerID,
		// Exit: false,
	}
	return cli.CheckpointCreate(context.Background(), containerID, opts)
}

// CheckpointExists check if a container checkpoint exists
func CheckpointExists(containerID string) ([]types.Checkpoint, error) {
	log.Debugf("Kill container %s", containerID)
	cli, err := getClient()
	if err != nil {
		return nil, err
	}
	opts := types.CheckpointListOptions{}
	return cli.CheckpointList(context.Background(), containerID, opts)
}
