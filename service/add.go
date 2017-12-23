package service

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/docker"
	"github.com/rs/xid"
)

func removeTar(archive string) {
	if err := os.RemoveAll(archive); err != nil {
		log.Printf("Failed to remove %s: %s", archive, err.Error())
	}
}

//Add a function
func Add(req *api.AddRequest) (*api.AddResponse, error) {

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
	imageName := "mufaas-" + name

	imageInfo, err := docker.ImageBuild(imageName, archive)
	if err != nil {
		return nil, err
	}

	return &api.AddResponse{
		Info: &api.FunctionInfo{
			Name: name,
			ID:   imageInfo.ID,
		},
	}, nil
}
