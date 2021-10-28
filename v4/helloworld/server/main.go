package main

import (
	"context"

	"github.com/asim/go-micro/plugins/server/grpc/v4"
	pb "github.com/xpunch/go-micro-example/v4/helloworld/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

type HelloWorld struct{}

func (h *HelloWorld) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	srv := micro.NewService(micro.Server(grpc.NewServer()), micro.Name("helloworld"))
	srv.Init()
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(HelloWorld)); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
