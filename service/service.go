package service

import (
	"errors"
	"net"

	"golang.org/x/net/context"

	"github.com/muka/mufaas/api"
	"google.golang.org/grpc"
)

var server *grpc.Server

type mu struct{}

func newMufaasService() api.MufaasServiceServer {
	return new(mu)
}

//Start the gRPC server
func Start(uri string) error {

	if uri == "" {
		return errors.New("gRPC uri not provided")
	}

	listen, err := net.Listen("tcp", uri)
	if err != nil {
		return err
	}
	server = grpc.NewServer()
	api.RegisterMufaasServiceServer(server, newMufaasService())

	return server.Serve(listen)
}

//Stop the gRPC server
func Stop() {
	if server == nil {
		return
	}
	server.Stop()
	server = nil
}

func (f *mu) Add(ctx context.Context, msg *api.AddRequest) (*api.AddResponse, error) {
	return Add(msg)
}

func (f *mu) Remove(ctx context.Context, msg *api.RemoveRequest) (*api.RemoveResponse, error) {
	return Remove(msg)
}

func (f *mu) List(ctx context.Context, msg *api.ListRequest) (*api.ListResponse, error) {
	return List(msg)
}

func (f *mu) Run(ctx context.Context, msg *api.RunRequest) (*api.RunResponse, error) {
	return Run(msg)
}
