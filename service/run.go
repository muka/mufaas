package service

import (
	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/docker"
)

//Run execute a function
func Run(req *api.RunRequest) (*api.RunResponse, error) {

	opts := docker.ExecOptions{
		Name:  req.Name,
		Args:  req.Args,
		Stdin: req.Stdin,
		Env:   req.Env,
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
