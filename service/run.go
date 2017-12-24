package service

import (
	"errors"
	"fmt"

	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/docker"
)

func Run(req *api.RunRequest) (*api.RunResponse, error) {

	if len(req.Name) == 0 {
		return nil, errors.New("Image name is empty")
	}

	images, err := docker.ImageList([]string{"reference=" + req.Name})
	if err != nil {
		return nil, err
	}

	if len(images) != 1 {
		return nil, fmt.Errorf("Image %s not found", req.Name)
	}

	opts := docker.ExecOptions{
		ImageName: req.Name,
		Args:      req.Args,
		Stdin:     req.Stdin,
		Env:       req.Env,
	}
	execResult, err := docker.Exec(opts)
	if err != nil {
		return nil, err
	}

	var runErr []byte
	var runOut []byte
	if execResult.Stderr != nil {
		runErr = execResult.Stderr.Bytes()
	}
	if execResult.Stdout != nil {
		runOut = execResult.Stdout.Bytes()
	}

	return &api.RunResponse{
		Err:    runErr,
		Output: runOut,
	}, nil
}
