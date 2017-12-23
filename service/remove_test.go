package service

import (
	"context"
	"errors"
	"testing"

	"github.com/muka/mufaas/api"
	log "github.com/sirupsen/logrus"
)

func containsFunction(client api.MufaasServiceClient, ids ...string) (bool, int, error) {

	ctx := context.Background()
	listReq := &api.ListRequest{}
	listRes, err := client.List(ctx, listReq)
	if err != nil {
		return false, 0, err
	}

	log.Debugf("Found %d images", len(listRes.Functions))

	var found int
	for _, fn := range listRes.Functions {
		for _, id := range ids {
			if fn.ID == id {
				found++
			}
		}
	}
	if found != len(ids) {
		return false, found, errors.New("Image not found")
	}
	return true, found, nil
}

func TestRemoveAll(t *testing.T) {

	runServer(t)

	client, conn, err := api.NewClient(grpcEndpoint)
	defer stopServer(conn)

	if err != nil {
		t.Fatal(err)
	}

	f1, err := createFunction(client)
	if err != nil {
		t.Fatalf("add failed: %s\n", err.Error())
	}

	if _, _, err := containsFunction(client, f1.ID); err != nil {
		t.Fatalf(err.Error())
	}

	if _, err := removeFunction(client); err != nil {
		t.Fatalf("Failed to remove functions: %s\n", err.Error())
	}

	if _, _, err := containsFunction(client); err != nil {
		t.Fatalf(err.Error())
	}

}
