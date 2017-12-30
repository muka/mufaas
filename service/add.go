package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/docker"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func removeTar(archive string) {
	if err := os.RemoveAll(archive); err != nil {
		log.Printf("Failed to remove %s: %s", archive, err.Error())
	}
}

//Add a function
func Add(req *api.AddRequest) (*api.AddResponse, error) {

	if req.Image == "" {

		if len(req.Source) == 0 {
			return nil, errors.New("Image source missing")
		}

		guid := xid.New().String()
		archive := path.Join(os.TempDir(), "mufaas_"+guid+".tar")

		err := ioutil.WriteFile(archive, req.Source, 0x700)
		if err != nil {
			return nil, err
		}
		defer removeTar(archive)

		name := req.Info.Name
		if name == "" {
			name = guid
		}
		imageName := name

		_, err = docker.ImageBuild(docker.ImageBuildOptions{
			Name:    imageName,
			Archive: archive,
		})
		if err != nil {
			return nil, err
		}

		req.Image = imageName

	} else {

		imgs, err := docker.ImageList([]string{"reference=" + req.Image})
		if err != nil {
			return nil, err
		}

		if len(imgs) != 1 {
			return nil, fmt.Errorf("Image %s not found", req.Image)
		}

	}

	container, err := docker.Create(&docker.CreateOptions{
		Cmd:        req.Info.Cmd,
		Env:        req.Info.Env,
		Image:      req.Image,
		Name:       req.Info.Name,
		Privileged: req.Info.Privileged,
	})
	if err != nil {
		return nil, err
	}

	return &api.AddResponse{
		Info: &api.FunctionInfo{
			Name: container.Name,
			ID:   container.ID,
		},
	}, nil
}
