package docker

import (
	"context"
	"fmt"

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
	Error   error
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

				// fmt.Printf("Event recieved: %s %s ", event.Action, event.Type)
				if event.Actor.Attributes != nil {

					name := event.Actor.Attributes["name"]
					switch event.Action {
					case "start":
						// fmt.Printf("Container started %s", name)
						break
					case "die":
						// fmt.Printf("Container exited %s", name)
						break
					}

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
				fmt.Printf("Error event recieved: %s", err.Error())
				ev := ContainerEvent{
					Error: err,
				}
				eventsChannel <- ev
			}
		}
	}
}
