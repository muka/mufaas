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

	forceRemove := req.Force

	filter := []string{}
	for _, name := range req.Name {
		filter = append(filter, "name="+name)
	}

	containers, err := docker.List(filter)
	if err != nil {
		return nil, err
	}

	log.Debugf("Got %d containers to remove", len(containers))
	for _, container := range containers {

		name := container.Names[0]
		r := &api.FunctionInfo{
			ID:   container.ID,
			Name: name,
		}
		log.Debugf("Remove container %s", container.ID)
		err := docker.Remove(container.ID, forceRemove)
		if err != nil {
			r.Error = err.Error()
		}
		res.Functions = append(res.Functions, r)
	}

	// remove dangling
	filter = []string{
		"label=" + docker.DefaultLabel + "=1", // only if managed by us
		"dangling=true",
	}
	images, err := docker.ImageList(filter)
	if err != nil {
		return nil, err
	}
	for _, image := range images {
		log.Debugf("Remove dangling image %s", image.ID)
		err := docker.ImageRemove(image.ID, true)
		if err != nil {
			log.Warnf("Failed to remove %s: %s", image.ID, err.Error())
		}
	}

	return res, nil
}
