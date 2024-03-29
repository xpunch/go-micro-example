package main

import (
	"context"

	_ "github.com/go-micro/plugins/v4/registry/etcd"
	_ "github.com/go-micro/plugins/v4/server/grpc"
	pb "github.com/xpunch/go-micro-example/v4/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

type HelloWorld struct{}

func (h *HelloWorld) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	if name, ok := metadata.Get(ctx, "name"); ok {
		rsp.Message = "Hello " + name
	} else {
		rsp.Message = "Hello " + req.Name
	}
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("helloworld.srv"),
	)
	srv.Init()
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(HelloWorld)); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
