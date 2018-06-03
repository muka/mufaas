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

	t1 := time.Now()

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

	t2 := time.Since(t1)
	log.Debugf("GetByName took %dms", t2.Nanoseconds()/1000000)

	containerID := container.ID

	if container.State != "running" {
		log.Debug("Container not running, starting")
		_, err = Start(ContainerStartOptions{
			ImageName: container.Image,
			Name:      containerName,
		})
		if err != nil {
			return nil, err
		}
		t3 := time.Since(t1)
		log.Debugf("ContainerStart took %dms", t3.Nanoseconds()/1000000)
	}

	ins, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return nil, err
	}

	t4 := time.Since(t1)
	log.Debugf("Container inspect took %dms", t4.Nanoseconds()/1000000)

	var cmd []string
	for _, e := range ins.Config.Env {
		p := strings.Split(e, "=")
		if p[0] == CmdEnvKey {
			decoded, err1 := base64.RawStdEncoding.DecodeString(p[1])
			if err1 != nil {
				return nil, err1
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

	wait := make(chan error, 1)
	exited := false

	// io.Writer
	outBuffer := bytes.NewBuffer([]byte{})
	errBuffer := bytes.NewBuffer([]byte{})
	go func() {
		_, berr := stdcopy.StdCopy(outBuffer, errBuffer, attachResp.Reader)
		if berr != nil {
			log.Debugf("Fail stdout copy: %s", berr.Error())
		}
		wait <- berr
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
		Tty:    false,
		Detach: false,
	})
	if err != nil {
		return nil, err
	}

	t5 := time.Since(t1)
	log.Debugf("ContainerExecStart took %dms", t5.Nanoseconds()/1000000)

	// negative timeout means no timeout
	if opts.Timeout > -1 {
		go func() {
			d := time.Second * time.Duration(opts.Timeout)
			time.Sleep(d)
			if !exited {
				log.Debugf("Forcing container exit by kill")
				kerr := Kill(containerID)
				if kerr != nil {
					log.Debugf("Error on kill: %s", kerr.Error())
				}
				wait <- kerr
			}
		}()
	}

	go func() {
		for {
			select {
			case ev := <-eventsChannel:
				if ev.Action == "quit_ch" {
					return
				}
				log.Debugf("Event %s %s", ev.Action, ev.ID)
				if ev.ID == containerID {
					if ev.Action == "die" {
						wait <- nil
						return
					}
				}
			}
		}
	}()

	// sleep until something happens
	err = <-wait
	if err != nil {
		return nil, err
	}

	exited = true
	eventsChannel <- ContainerEvent{Action: "quit_ch"}

	if opts.Remove {
		err = Remove(containerID, true)
		if err != nil {
			log.Debugf("Remove err: %s", err.Error())
		}
	}

	t6 := time.Since(t1)
	log.Debugf("Exec took %dms", t6.Nanoseconds()/1000000)

	return &ExecResult{
		Stdout: outBuffer,
		Stderr: errBuffer,
	}, nil
}
