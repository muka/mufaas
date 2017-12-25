package service

import (
	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/docker"
	log "github.com/sirupsen/logrus"
)

//Remove handle a function removal
func Remove(req *api.RemoveRequest) (*api.RemoveResponse, error) {

	res := &api.RemoveResponse{
		Functions: []*api.FunctionInfo{},
	}

	filter := []string{}
	for _, name := range req.Name {
		filter = append(filter, "reference=mufaas-"+name)
	}

	images, err := docker.ImageList(filter)
	if err != nil {
		return nil, err
	}

	for _, image := range images {
		r := &api.FunctionInfo{
			ID: image.ID,
		}
		log.Debugf("Remove image %s", image.ID)
		err := docker.ImageRemove(image.ID)
		if err != nil {
			r.Error = err.Error()
		}
		res.Functions = append(res.Functions, r)
	}

	return res, nil
}
