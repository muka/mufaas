package service

import (
	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/docker"
)

//Remove handle a function removal
func Remove(req *api.RemoveRequest) (*api.RemoveResponse, error) {

	res := &api.RemoveResponse{
		Functions: []*api.FunctionInfo{},
	}

	images, err := docker.ImageList(req.Filter)
	if err != nil {
		return nil, err
	}

	for _, image := range images {
		r := &api.FunctionInfo{
			ID: image.ID,
		}
		err := docker.ImageRemove(image.ID)
		if err != nil {
			r.Error = err.Error()
		}
		res.Functions = append(res.Functions, r)
	}

	return res, nil
}
