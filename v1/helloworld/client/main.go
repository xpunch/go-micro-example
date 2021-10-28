package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	"github.com/xpunch/go-micro-example/v1/helloworld/proto"
)

func main() {
	srv := micro.NewService(micro.Name("client"))
	srv.Init()

	greeter := proto.NewHelloworldService("helloworld", srv.Client())
	resp, err := greeter.Call(context.TODO(), &proto.Request{Name: "Client"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)
}
