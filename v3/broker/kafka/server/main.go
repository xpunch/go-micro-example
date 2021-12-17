package main

import (
	"context"

	_ "github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
)

type Message struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"ts"`
}

func main() {
	srv := micro.NewService(
		micro.Name("kafka-server"),
	)
	srv.Init()
	if err := srv.Server().Subscribe(srv.Server().NewSubscriber("kafka-topic", subscribe)); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func subscribe(ctx context.Context, msg *Message) error {
	logger.Info(msg)
	return nil
}
