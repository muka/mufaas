package service

import (
	"context"
	"os"
	"testing"
	"time"

	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/muka/mufaas/docker"
	"google.golang.org/grpc"

	"github.com/muka/mufaas/api"
	"github.com/muka/mufaas/util"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

const grpcEndpoint = "localhost:5001"

func TestMain(m *testing.M) {

	var v bool
	for _, arg := range os.Args {
		if arg == "-test.v=true" {
			v = true
			break
		}
	}

	if v {
		log.SetLevel(log.DebugLevel)
	}

	f := []string{"label=mufaas=1"}
	list, err := docker.ImageList(f)

	if err != nil {
		panic(err)
	}

	for _, image := range list {
		docker.ImageRemove(image.ID, true)
	}

	log.Debugf("Removed previous images")

	os.Exit(m.Run())
}

func runServer(t *testing.T) {
	go func() {
		err := Start(grpcEndpoint)
		if err != nil {
			t.Fatal(err)
		}
	}()
	//wait for the service to start
	time.Sleep((time.Millisecond * 500))
}

func stopServer(conn *grpc.ClientConn) {
	conn.Close()
	Stop()
	//wait for the service to start
	time.Sleep((time.Millisecond * 500))
}

func TestServer(t *testing.T) {

	runServer(t)

	client, conn, err := api.NewClient(grpcEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	defer stopServer(conn)

	ctx := context.Background()
	req := &api.ListRequest{}
	_, err = client.List(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

}

func removeFunction(client api.MufaasServiceClient, force bool, names ...string) (bool, error) {

	log.Debugf("removeFunction: Removing images: %s", names)

	ctx := context.Background()

	filter := []string{}
	for _, n := range names {
		filter = append(filter, "name="+n)
	}
	list, err := client.List(ctx, &api.ListRequest{Filter: filter})
	if err != nil {
		return false, err
	}

	// log.Debugf("******** %++v", list)
	listLen := len(list.Functions)

	log.Debugf("removeFunction: Found %d unique image ID", listLen)
	// log.Debugf("removeFunction: %+v", imageIDs)

	rmReq := &api.RemoveRequest{Name: names, Force: true}
	rmRes, err := client.Remove(ctx, rmReq)
	if err != nil {
		return false, err
	}

	rmLen := len(rmRes.Functions)
	log.Debugf("removeFunction: Remove reported %d uniqe images", rmLen)
	// log.Debugf("-------- %+v", rmRes.Functions)

	//TODO image list is inconsistent, need more work over that
	if rmLen != listLen {
		log.Errorf("removeFunction: Removed functions count not matching (rm %d = ids %d)", rmLen, listLen)
		return false, fmt.Errorf("removeFunction: Removed functions count not matching (rm %d = ids %d)", rmLen, listLen)
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
