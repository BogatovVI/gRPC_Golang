package app

import (
	"context"
	"grpc_golang/internal/service"
	greeter "grpc_golang/pb/proto"

	"google.golang.org/grpc"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listener := service.MustCreateListener(ctx)

	grpcServer := grpc.NewServer(grpc.EmptyServerOption{})

	greeterGrpc := &service.GreeterGrpc{}

	greeter.RegisterGreeterServer(grpcServer, greeterGrpc)
	grpcServer.Serve(listener)
}
