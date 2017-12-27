package docker

import (
	"strings"
	"testing"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func TestExecNoArgs(t *testing.T) {

	uniqid := xid.New().String()
	imageID := doBuild(t, "../test/hello", "mufaas/hello-"+uniqid)

	log.Debugf("Created image %s", imageID)

	opts := ExecOptions{
		Name:      "exec_noargs_" + uniqid,
		ImageName: "mufaas/hello-" + uniqid,
	}

	res, err := Exec(opts)
	if err != nil {
		t.Fatalf("Exec failed: %s", err.Error())
	}

	if len(res.Stdout.String()) == 0 {
		t.Fatal("Unexpected empty output")
	}

	err = ImageRemove(imageID, true)
	if err != nil {
		t.Fatal(err)
	}

	log.Debugf("Out: \n\n%s", res.Stdout.String())
}

func TestExecWithArgs(t *testing.T) {

	TestBuild(t)

	hello := "world"
	opts := ExecOptions{
		Name:      "exec_test_hello_args",
		ImageName: "mufaas/hello",
		Args:      []string{hello},
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
	if !strings.Contains(lines[0], hello) {
		t.Fatalf("Expecting to find `hello %s` at first line", hello)
	}

}

func TestExecWithTimeout(t *testing.T) {

	uniqid := xid.New().String()
	imageName := "mufaas/test-timeout-" + uniqid
	doBuild(t, "../test/timeout", imageName)

	opts := ExecOptions{
		Name:      "exec_test_timeout_" + uniqid,
		ImageName: imageName,
		Args:      []string{"timeout"},
		Timeout:   2,
	}

	res, err := Exec(opts)
	if err != nil {
		t.Fatalf("Exec failed: %s", err.Error())
	}

	if len(res.Stdout.String()) == 0 {
		t.Fatal("Unexpected empty output")
	}

	log.Debugf("Out: \n\n%s", res.Stdout.String())

}
