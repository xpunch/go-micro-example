package main

import (
	"context"

	_ "github.com/go-micro/plugins/v4/broker/kafka"
	_ "github.com/go-micro/plugins/v4/broker/mqtt"
	_ "github.com/go-micro/plugins/v4/registry/etcd"
	pb "github.com/xpunch/go-micro-example/v4/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

type Message struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type subscriber struct{}

func (s *subscriber) Subscribe(ctx context.Context, msg *pb.Message) error {
	logger.Info(msg)
	return nil
}

func subscribe(ctx context.Context, msg *Message) error {
	logger.Info(msg)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("broker-server"),
	)
	srv.Init()
	if err := micro.RegisterSubscriber("micro.topic.json", srv.Server(), subscribe); err != nil {
		logger.Fatal(err)
	}
	if err := micro.RegisterSubscriber("micro.topic.protobuf", srv.Server(), &subscriber{}); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
