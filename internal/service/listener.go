package service

import (
	"context"
	"fmt"
	"net"
)

func MustCreateListener(ctx context.Context) net.Listener {
	listenConfig := net.ListenConfig{
		KeepAlive: -1,
		KeepAliveConfig: net.KeepAliveConfig{
			Enable: true,
		},
	}

	listener, err := listenConfig.Listen(ctx, "tcp", ":8084")

	if err != nil {
		panic(fmt.Sprintf("Error create listener for grpc service: %s", err.Error()))
	}

	return listener
}
