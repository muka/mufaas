package docker

import (
	"log"
	"strings"
	"testing"

	"github.com/rs/xid"
)

func TestExecNoArgs(t *testing.T) {

	TestBuild(t)

	opts := ExecOptions{
		Name:      "exec_test_hello",
		ImageName: "mufaas/hello",
	}

	res, err := Exec(opts)
	if err != nil {
		t.Fatalf("Exec failed: %s", err.Error())
	}

	if len(res.Stdout.String()) == 0 {
		t.Fatal("Unexpected empty output")
	}

	log.Printf("Out: \n\n%s", res.Stdout.String())
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

	log.Printf("Out: \n\n%s", res.Stdout.String())

	lines := strings.Split(res.Stdout.String(), "\n")
	if !strings.Contains(lines[0], hello) {
		t.Fatalf("Expecting to find `hello %s` at first line", hello)
	}

}

func TestExecWithTimeout(t *testing.T) {

	imageName := "mufaas/test-timeout-" + xid.New().String()
	doBuild(t, "../test/timeout", imageName)

	opts := ExecOptions{
		Name:      "exec_test_timeout",
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

	log.Printf("Out: \n\n%s", res.Stdout.String())

}
