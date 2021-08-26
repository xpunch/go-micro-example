package main

import (
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/xpunch/go-micro-example/v2/helloworld/proto"
)

func main() {
	srv := micro.NewService(micro.Name("client"))
	srv.Init()

	greeter := proto.NewHelloworldService("helloworld", srv.Client())
	resp, err := greeter.Call(context.TODO(), &proto.Request{Name: "Client"})
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(resp)
}
