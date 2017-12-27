package docker

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
)

//ContainerEvent store a container event
type ContainerEvent struct {
	ID      string
	Name    string
	Action  string
	Message events.Message
}

var eventsChannel = make(chan ContainerEvent)

func getEventsChannel() chan ContainerEvent {
	return eventsChannel
}

func watchEvents() error {
	cli, err := getClient()
	if err != nil {
		return err
	}
	ctx := context.Background()
	f := filters.NewArgs()
	f.Add("label", DefaultLabel+"=1")

	msgChan, errChan := cli.Events(ctx, types.EventsOptions{
		Filters: f,
	})

	for {
		select {
		case event := <-msgChan:
			if &event != nil {

				if event.Actor.Attributes != nil {

					name := event.Actor.Attributes["name"]
					ev := ContainerEvent{
						Action:  event.Action,
						ID:      event.ID,
						Name:    name,
						Message: event,
					}
					eventsChannel <- ev

				}
			}
		case err := <-errChan:
			if err != nil {
				log.Warnf("Error event recieved: %s", err.Error())
			}
		}
	}
}
