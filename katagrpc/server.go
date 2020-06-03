package katagrpc

import (
	"github.com/farislr/core-proto/kata"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

func Init(server *grpc.Server, logger log.Logger) {
	var service kata.StringServiceServer
	{
		service = NewService(logger)
	}

	kata.RegisterStringServiceServer(server, service)
}
