package main

import (
	"sync"

	mgrpc "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	pb "github.com/xpunch/go-micro-example/v3/event/proto"
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
