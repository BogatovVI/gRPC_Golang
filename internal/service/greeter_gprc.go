package service

import (
	"context"
	"fmt"
	greeter "grpc_golang/pb/proto"
)

type GreeterGrpc struct {
	greeter.UnimplementedGreeterServer
}

func (gg *GreeterGrpc) SayHello(
	ctx context.Context, helloRequest *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{
		Message: fmt.Sprintf("Hello %s", helloRequest.Name),
	}, nil
}
