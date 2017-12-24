package service

import (
	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/docker"
)

//List handles a functions listing request
func List(req *api.ListRequest) (*api.ListResponse, error) {
	images, err := docker.ImageList(req.Filter)
	if err != nil {
		return nil, err
	}
	var list []*api.FunctionInfo
	for _, image := range images {
		item := &api.FunctionInfo{
			ID:   image.ID,
			Name: image.RepoTags[0],
			Type: image.Labels["mufaas-type"],
		}
		list = append(list, item)
	}
	res := &api.ListResponse{
		Functions: list,
	}
	return res, nil
}
