package main

import (
	"sync"

	mgrpc "github.com/go-micro/plugins/v4/server/grpc"
	pb "github.com/xpunch/go-micro-example/v4/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	AccessCount   int64
	AccessMethods map[string]int64
	mu            sync.RWMutex
)

func main() {
	srv := micro.NewService(
		micro.Server(mgrpc.NewServer()),
		micro.Name("statistics"),
		micro.Address(":8082"),
	)
	srv.Init()
	hdl := &handler{}
	if err := pb.RegisterStatisticsServiceHandler(srv.Server(), hdl); err != nil {
		logger.Fatal(err)
	}
	sub := &subscriber{}
	if err := srv.Server().Subscribe(srv.Server().NewSubscriber("accesslogs", sub.AccessLog)); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func init() {
	AccessMethods = make(map[string]int64)
}
