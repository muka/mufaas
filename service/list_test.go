package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/util"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func removeFunction(client api.MufaasServiceClient, force bool, names ...string) (bool, error) {

	log.Debugf("removeFunction: Removing images: %s", names)

	ctx := context.Background()
	filter := []string{}
	if !force {
		filter = append(filter, "dangling=true")
	}
	for _, n := range names {
		filter = append(filter, "reference="+n)
	}
	list, err := client.List(ctx, &api.ListRequest{Filter: filter})
	if err != nil {
		return false, err
	}

	var imageIDs []string
	for _, img := range list.Functions {
		var exists bool
		for _, imgid := range imageIDs {
			if imgid == img.ID {
				exists = true
			}
		}
		if !exists {
			imageIDs = append(imageIDs, img.ID)
		}
	}

	log.Debugf("removeFunction: Found %d unique image ID", len(imageIDs))

	rmReq := &api.RemoveRequest{Name: names, Force: true}
	rmRes, err := client.Remove(ctx, rmReq)
	if err != nil {
		return false, err
	}

	log.Debugf("removeFunction: Remove reported %d images", len(rmRes.Functions))

	if len(rmRes.Functions) != len(imageIDs) {
		return false, fmt.Errorf("removeFunction: Removed functions count not matching (rm %d = ids %d)", len(rmRes.Functions), len(imageIDs))
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

	// add tmp file to trick docker build and create a new image
	tmpFile := filepath.Join(dir, xid.New().String()+".txt")
	err := ioutil.WriteFile(tmpFile, []byte{}, 0x755)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile)

	err = util.CreateArchive(dir)
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
			Name: "test_" + xid.New().String(),
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
