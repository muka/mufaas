package docker

import (
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestExecNoArgs(t *testing.T) {

	info := createContainer(t, "hello")
	opts := ExecOptions{Name: info.Name}

	res, err := Exec(opts)
	if err != nil {
		t.Fatalf("Exec failed: %s", err.Error())
	}

	if len(res.Stdout.String()) == 0 {
		t.Fatal("Unexpected empty output")
	}

	removeContainer(t, info.ID)
	removeImage(t, info.ImageID)

	log.Debugf("Out: \n\n%s", res.Stdout.String())
}

func TestExecWithArgs(t *testing.T) {

	info := createContainer(t, "hello")
	opts := ExecOptions{
		Name: info.Name,
		Args: []string{info.ImageName},
	}

	res, err := Exec(opts)
	if err != nil {
		t.Fatalf("Exec failed: %s", err.Error())
	}

	if len(res.Stdout.String()) == 0 {
		t.Fatal("Unexpected empty output")
	}

	log.Debugf("Out: \n\n%s", res.Stdout.String())

	lines := strings.Split(res.Stdout.String(), "\n")
	if !strings.Contains(lines[0], info.ImageName) {
		t.Fatalf("Expecting to find `hello %s` at first line", info.ImageName)
	}

	removeContainer(t, info.ID)
	removeImage(t, info.ImageID)

}

func TestExecWithTimeout(t *testing.T) {

	info := createContainer(t, "timeout")
	opts := ExecOptions{
		Name:    info.Name,
		Args:    []string{"timeout"},
		Timeout: 2,
	}

	res, err := Exec(opts)
	if err != nil {
		t.Fatalf("Exec failed: %s", err.Error())
	}

	if len(res.Stdout.String()) == 0 {
		t.Fatal("Unexpected empty output")
	}

	removeContainer(t, info.ID)
	removeImage(t, info.ImageID)

	log.Debugf("Out: \n\n%s", res.Stdout.String())

}
