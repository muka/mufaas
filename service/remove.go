package service

import (
	"strings"

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

	filter := []string{
		"label=" + docker.DefaultLabel + "=1", // only if managed by us
	}

	// if !forceRemove { // only unused by container, otherwise an issue may arise
	// 	filter = append(filter, "dangling=true")
	// }

	for _, name := range req.Name {
		filter = append(filter, "reference="+name)
	}

	images, err := docker.ImageList(filter)
	if err != nil {
		return nil, err
	}

	log.Debugf("Got %d images to remove", len(images))
	for _, image := range images {

		name := image.RepoTags[0]
		name = name[:strings.Index(name, ":")]

		r := &api.FunctionInfo{
			ID:   image.ID,
			Name: name,
		}
		log.Debugf("Remove image %s", image.ID)
		err := docker.ImageRemove(image.ID, forceRemove)
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
	images, err = docker.ImageList(filter)
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
