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

	log.Debugf("Got %d functions to remove", len(containers))

	imagesList := map[string]bool{}
	for _, container := range containers {

		name := container.Names[0]
		r := &api.FunctionInfo{
			ID:   container.ID,
			Name: name,
		}

		if container.State == "running" {
			err1 := docker.Kill(container.ID)
			if err1 != nil {
				log.Warnf("Failed to kill container: %s", err1.Error())
			}
		}

		log.Debugf("Remove container %s", container.ID)
		err1 := docker.Remove(container.ID, forceRemove)
		if err1 != nil {
			r.Error = err1.Error()
		}
		res.Functions = append(res.Functions, r)

		imagesList[container.Image] = true
		imagesList[container.ImageID] = true
	}

	for imageID := range imagesList {

		containerList, err1 := docker.List([]string{
			"ancestor=" + imageID,
		})
		if err1 != nil {
			return nil, err1
		}

		if len(containerList) > 0 {
			var cnames string
			for _, c := range containerList {
				cnames += " " + c.Names[0]
			}
			log.Debugf("Skip image %s as used by %s", imageID, cnames)
			continue
		}

		err = docker.ImageRemove(imageID, true)
		if err != nil {
			log.Warnf("Failed to remove %s: %s", imageID, err.Error())
		}
	}

	// remove dangling images
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
