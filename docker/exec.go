package docker

import (
	"bytes"
	"errors"
	"io"
	"time"

	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	log "github.com/sirupsen/logrus"
)

var defaultTimeout int64 = 3

//ExecOptions control how a container is executed
type ExecOptions struct {
	Name  string
	Cmd   []string
	Env   []string
	Stdin []byte
	Args  []string
	// Timeout in second to stop the container
	Timeout int64
	Remove  bool
}

//ExecResult return the execution results
type ExecResult struct {
	ID     string
	Stdout *bytes.Buffer
	Stderr *bytes.Buffer
}

// Exec spawn a container and wait for its output
func Exec(opts ExecOptions) (*ExecResult, error) {

	if len(opts.Name) == 0 {
		return nil, errors.New("Function name is empty")
	}

	cli, err := getClient()
	if err != nil {
		return nil, err
	}

	// set default TTL
	if opts.Timeout == 0 {
		opts.Timeout = defaultTimeout
	}

	ctx := context.Background()
	containerName := opts.Name

	containers, err := List([]string{"name=" + containerName})
	if err != nil {
		return nil, err
	}

	if len(containers) != 1 {
		return nil, fmt.Errorf("Function %s not found", containerName)
	}

	containerID := containers[0].ID

	startConfig := types.ContainerStartOptions{}
	if err = cli.ContainerStart(ctx, containerID, startConfig); err != nil {
		return nil, err
	}

	attachConfig := types.ContainerAttachOptions{
		Logs:   false,
		Stdin:  false,
		Stdout: true,
		Stderr: true,
		Stream: true,
	}
	conn, err := cli.ContainerAttach(ctx, containerID, attachConfig)
	if err != nil {
		return nil, err
	}

	log.Debugf("Started %s", containerName)

	wait := make(chan bool, 1)

	go func() {
		d := time.Second * time.Duration(opts.Timeout)
		time.Sleep(d)
		kerr := Kill(containerID)
		if kerr != nil {
			log.Debugf("Error on kill: %s", kerr.Error())
		}
		wait <- true
	}()

	go func() {
		ch := getEventsChannel()
		for {
			select {
			case ev := <-ch:
				if ev.ID == containerID {
					if ev.Action == "die" {
						wait <- true
						return
					}
				}
			}
		}
	}()

	// sleep until something happens
	<-wait

	var outBuffer bytes.Buffer
	_, berr := io.Copy(&outBuffer, conn.Reader)
	if berr != nil {
		if berr != io.EOF {
			log.Debugf("Fail stdout copy: %s\n", berr.Error())
		}
	}

	defer conn.Close()

	if opts.Remove {
		err = Remove(containerID, true)
		if err != nil {
			log.Debugf("Remove err: %s", err.Error())
		}
	}

	return &ExecResult{
		Stdout: &outBuffer,
		Stderr: nil,
	}, nil
}
