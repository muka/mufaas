package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/cli"
)

func removeFunction(client api.MufaasServiceClient, ids ...string) (bool, error) {

	filter := []string{}
	for _, id := range ids {
		filter = append(filter, "reference="+id)
	}

	ctx := context.Background()
	rmReq := &api.RemoveRequest{Filter: filter}
	rmRes, err := client.Remove(ctx, rmReq)

	if err != nil {
		return false, err
	}

	for _, f := range rmRes.Functions {
		if f.Error != "" {
			return false, fmt.Errorf("[%s] %s", f.ID, f.Error)
		}
	}

	return true, nil
}

func createFunction(client api.MufaasServiceClient) (*api.FunctionInfo, error) {

	dir := "../test/hello"

	err := cli.CreateArchive(dir)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(dir + ".tar")
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	addReq := &api.AddRequest{
		Info: &api.FunctionInfo{
			Type: "node",
			Name: "test1",
		},
		Source: content,
	}

	addRes, err := client.Add(ctx, addReq)
	if err != nil {
		return nil, err
	}

	if addRes.Info.Error != "" {
		return nil, fmt.Errorf("Add error: %s", addRes.Info.Error)
	}

	return addRes.Info, nil
}

func TestList(t *testing.T) {

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

	ctx := context.Background()
	listReq := &api.ListRequest{}
	listRes, err := client.List(ctx, listReq)
	if err != nil {
		t.Fatal(err)
	}

	var found bool
	var ids []string
	for _, f := range listRes.Functions {
		if f.ID == addf.ID {
			found = true
			ids = append(ids, f.ID)
		}
	}

	if !found {
		fmt.Printf("FAIL: function not found %s\n", addf.ID)
		t.Fail()
	}

	if _, err := removeFunction(client, ids...); err != nil {
		t.Fatalf("Failed to remove functions: %s\n", err.Error())
	}

}
