package main

import (
	"context"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	pb "github.com/xpunch/go-micro-example/v3/helloworld/proto"
)

type HelloWorld struct{}

func (h *HelloWorld) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	srv := micro.NewService(micro.Name("helloworld"))
	srv.Init()
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(HelloWorld)); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
