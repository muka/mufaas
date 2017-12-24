package service

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/muka/mufaas/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const grpcEndpoint = "localhost:5001"

func TestMain(m *testing.M) {

	var v = flag.Bool("v", false, "verbose")
	flag.Parse()
	if *v {
		log.SetLevel(log.DebugLevel)
	}

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
