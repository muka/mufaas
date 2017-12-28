package service

import (
	"context"
	"testing"

	"github.com/muka/mufaas/api"
	log "github.com/sirupsen/logrus"
)

func TestRunEmptyName(t *testing.T) {

	runServer(t)

	client, conn, err := api.NewClient(grpcEndpoint)
	defer stopServer(conn)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	req := &api.RunRequest{
		Name: "",
		Args: []string{},
	}
	_, err = client.Run(ctx, req)
	if err == nil {
		t.Fatal("Expected exception with empty name")
	}

}

func TestRunMissingFunction(t *testing.T) {

	runServer(t)

	client, conn, err := api.NewClient(grpcEndpoint)
	defer stopServer(conn)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	req := &api.RunRequest{
		Name: "non-existing",
		Args: []string{},
	}
	_, err = client.Run(ctx, req)
	if err == nil {
		t.Fatal("Expected exception with non existing image")
	}

}

func TestRun(t *testing.T) {

	runServer(t)

	client, conn, err := api.NewClient(grpcEndpoint)
	defer stopServer(conn)

	if err != nil {
		t.Fatal(err)
	}

	addf, err := createFunction(client)
	if err != nil {
		t.Fatalf("create fn failed: %s\n", err.Error())
	}

	log.Debugf("Created %s", addf.Name)
	ctx := context.Background()
	req := &api.RunRequest{
		Name: addf.Name,
		Args: []string{"run", "test"},
	}

	res, err := client.Run(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	log.Debugf("Out: %s", res.Output)

	if _, err := removeFunction(client, true, addf.Name); err != nil {
		t.Fatalf("Failed to remove functions: %s\n", err.Error())
	}

}
