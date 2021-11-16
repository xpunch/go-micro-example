package main

import (
	"context"

	"github.com/asim/go-micro/plugins/client/grpc/v4"
	_ "github.com/asim/go-micro/plugins/registry/etcd/v4"
	pb "github.com/xpunch/go-micro-example/v4/helloworld/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	srv := micro.NewService(micro.Client(grpc.NewClient()), micro.Name("client"))
	srv.Init()

	greeter := pb.NewHelloworldService("helloworld", srv.Client())
	resp, err := greeter.Call(context.TODO(), &pb.Request{Name: "Client"})
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(resp)
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
