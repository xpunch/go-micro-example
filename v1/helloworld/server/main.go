package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	"github.com/xpunch/go-micro-example/v1/helloworld/proto"
)

type HelloWorld struct{}

func (h *HelloWorld) Call(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	srv := micro.NewService(micro.Name("helloworld"))
	srv.Init()
	if err := proto.RegisterHelloworldHandler(srv.Server(), new(HelloWorld)); err != nil {
		log.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
