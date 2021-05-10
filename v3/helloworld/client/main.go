package main

import (
	"context"

	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	pb "github.com/xpunch/go-micro-examples/v3/helloworld/proto"
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
}
