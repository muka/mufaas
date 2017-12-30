package docker

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
	log "github.com/sirupsen/logrus"
)

var defaultTimeout int64 = 3

//ExecOptions control how a container is executed
type ExecOptions struct {
	Name       string
	Cmd        []string
	Env        []string
	Stdin      []byte
	Args       []string
	Privileged bool
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

	containerName := opts.Name
	ctx := context.Background()

	// set default TTL
	if opts.Timeout == 0 {
		opts.Timeout = defaultTimeout
	}

	container, err := GetByName(containerName)
	if err != nil {
		return nil, err
	}

	containerID := container.ID

	log.Debugf("Container status: %s", container.State)
	ins, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return nil, err
	}

	var cmd []string
	for _, e := range ins.Config.Env {
		p := strings.Split(e, "=")
		if p[0] == CmdEnvKey {
			decoded, err := base64.StdEncoding.DecodeString(p[1] + "=")
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(decoded, &cmd)
			if err != nil {
				return nil, err
			}
		}
	}

	if len(opts.Args) > 0 {
		cmd = append(cmd, opts.Args...)
	}

	execConfig := types.ExecConfig{
		Cmd:          cmd,
		Env:          opts.Env,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
		Privileged:   opts.Privileged,
	}

	r, err := cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		return nil, err
	}

	attachResp, err := cli.ContainerExecAttach(ctx, r.ID, execConfig)
	if err != nil {
		return nil, err
	}
	defer attachResp.Close()

	// io.Writer
	outBuffer := bytes.NewBuffer([]byte{})
	errBuffer := bytes.NewBuffer([]byte{})
	go func() {
		_, berr := stdcopy.StdCopy(outBuffer, errBuffer, attachResp.Reader)
		// _, berr := io.Copy(outBuffer, attachResp.Reader)
		if berr != nil {
			if berr == io.EOF {
				log.Debugf("GOT EOF %s", berr.Error())
			} else {
				log.Debugf("Fail stdout copy: %s", berr.Error())
			}
		}
	}()

	// for stdin, see cli/command/container/hijack.go in docker/cli
	// io.Writer
	if opts.Stdin != nil {
		inputBuffer := bytes.NewBuffer(opts.Stdin)
		_, err = io.Copy(attachResp.Conn, inputBuffer)
		if err != nil {
			return nil, err
		}
	}

	log.Debugf("Exec command %s", cmd)
	err = cli.ContainerExecStart(ctx, r.ID, types.ExecStartCheck{
		Tty: false,
	})
	if err != nil {
		return nil, err
	}

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
				log.Debugf("Event %s %s", ev.Action, ev.ID)
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
	log.Debug("Waiting for completion")
	<-wait

	if opts.Remove {
		err = Remove(containerID, true)
		if err != nil {
			log.Debugf("Remove err: %s", err.Error())
		}
	}

	return &ExecResult{
		Stdout: outBuffer,
		Stderr: errBuffer,
	}, nil
}
