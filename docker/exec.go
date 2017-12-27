package docker

import (
	"bytes"
	"io"
	"regexp"
	"time"

	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

var defaultTimeout int64 = 3

//ExecOptions control how a container is executed
type ExecOptions struct {
	Name      string
	Cmd       []string
	Env       []string
	Stdin     []byte
	Args      []string
	ImageName string
	// Timeout in second to stop the container
	Timeout int64
}

//ExecResult return the execution results
type ExecResult struct {
	ID     string
	Stdout *bytes.Buffer
	Stderr *bytes.Buffer
}

// Exec spawn a container and wait for its output
func Exec(opts ExecOptions) (*ExecResult, error) {
	cli, err := getClient()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	// set default TTL
	if opts.Timeout == 0 {
		opts.Timeout = defaultTimeout
	}

	containerName := xid.New().String()
	if opts.Name != "" {
		containerName = opts.Name
		reg, rerr := regexp.Compile("[^a-zA-Z0-9]+")
		if rerr != nil {
			return nil, rerr
		}
		containerName = reg.ReplaceAllString(containerName, "")
	}

	containerName = fmt.Sprintf("mufaas-%s", containerName)

	var cmd []string
	if len(opts.Cmd) == 0 {
		imgfilter := filters.NewArgs()
		imgfilter.Add("reference", opts.ImageName)
		listOptions := types.ImageListOptions{
			All:     true,
			Filters: imgfilter,
		}
		var imageID string
		imageList, ilerr := cli.ImageList(ctx, listOptions)
		if ilerr != nil {
			return nil, ilerr
		}
		if len(imageList) == 1 {
			imageID = imageList[0].ID
		} else {
			return nil, fmt.Errorf("Image not found %s", opts.ImageName)
		}

		imageInfo, _, iierr := cli.ImageInspectWithRaw(ctx, imageID)
		if iierr != nil {
			return nil, iierr
		}

		cmd = append(imageInfo.Config.Cmd, opts.Args...)
	} else {
		cmd = append(opts.Cmd, opts.Args...)
	}

	log.Debugf("Creating container %s (from %s)\n", containerName, opts.ImageName)
	containerConfig := &container.Config{
		Cmd:          cmd,
		Env:          opts.Env,
		Image:        opts.ImageName,
		AttachStdin:  false,
		AttachStderr: true,
		AttachStdout: true,
		Tty:          true,
		StdinOnce:    true,
		Labels: map[string]string{
			DefaultLabel: "1",
		},
	}

	hostConfig := &container.HostConfig{
		AutoRemove: false,
	}
	netConfig := &network.NetworkingConfig{}
	resp, cerr := cli.ContainerCreate(ctx, containerConfig, hostConfig, netConfig, containerName)
	if cerr != nil {
		return nil, cerr
	}

	containerID := resp.ID
	log.Debugf("Created container %s", containerID)

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

	log.Debugf("Started %s\n", containerID)

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

	err = Remove(containerID)
	if err != nil {
		log.Debugf("Remove err: %s", err.Error())
	}

	return &ExecResult{
		Stdout: &outBuffer,
		Stderr: nil,
	}, nil
}
