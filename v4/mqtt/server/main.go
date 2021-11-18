package main

import (
	"context"

	"github.com/asim/go-micro/plugins/broker/mqtt/v4"
	pb "github.com/xpunch/go-micro-example/v4/mqtt/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

type Message struct {
	ID        string `json:"id"`
	Message   string `jsno:"msg"`
	Timestamp int64  `json:"ts"`
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
		micro.Name("mqtt-server"),
		micro.Broker(mqtt.NewBroker()),
	)
	srv.Init()
	if err := micro.RegisterSubscriber("micro.broker.mqtt.json", srv.Server(), subscribe); err != nil {
		logger.Fatal(err)
	}
	if err := micro.RegisterSubscriber("micro.broker.mqtt.protobuf", srv.Server(), &subscriber{}); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
