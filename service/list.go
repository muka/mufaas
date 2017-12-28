package service

import (
	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/docker"
)

//List handles a functions listing request
func List(req *api.ListRequest) (*api.ListResponse, error) {
	containers, err := docker.List(req.Filter)
	if err != nil {
		return nil, err
	}
	var list []*api.FunctionInfo
	for _, container := range containers {
		item := &api.FunctionInfo{
			ID:   container.ID,
			Name: container.Names[0],
			Type: container.Labels["mufaas-type"],
		}
		list = append(list, item)
	}
	res := &api.ListResponse{
		Functions: list,
	}
	return res, nil
}
