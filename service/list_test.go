package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/muka/mufaas/api"
)

func TestList(t *testing.T) {

	runServer(t)

	client, conn, err := api.NewClient(grpcEndpoint)
	defer stopServer(conn)

	if err != nil {
		t.Fatal(err)
	}

	addf, err := createFunction(client, "TestList")
	if err != nil {
		t.Fatalf("create fn failed: %s\n", err.Error())
	}

	ctx := context.Background()
	listReq := &api.ListRequest{}
	listRes, err := client.List(ctx, listReq)
	if err != nil {
		t.Fatal(err)
	}

	var found bool
	for _, f := range listRes.Functions {
		if f.ID == addf.ID {
			found = true
		}
	}

	if !found {
		fmt.Printf("FAIL: function not found %s\n", addf.ID)
		t.Fail()
	}

	if _, err := removeFunction(client, true, addf.Name); err != nil {
		t.Fatalf("Failed to remove functions: %s\n", err.Error())
	}

}
