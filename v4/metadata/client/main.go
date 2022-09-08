package main

import (
	"context"

	_ "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"
	pb "github.com/xpunch/go-micro-example/v4/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

func main() {
	srv := micro.NewService(
		micro.Name("helloworld.cli"),
	)
	srv.Init()

	greeter := pb.NewHelloworldService("helloworld.srv", srv.Client())
	ctx := metadata.Set(context.TODO(), "name", "World")
	resp, err := greeter.Call(ctx, &pb.Request{})
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(resp)
}
