package main

import (
	"context"
	"time"

	_ "github.com/asim/go-micro/plugins/broker/kafka/v3"
	mgrpc "github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/logger"
	"github.com/google/uuid"
)

type Message struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"ts"`
}

func main() {
	srv := micro.NewService(
		micro.Name("kafka-client"),
		micro.Client(mgrpc.NewClient()),
	)
	srv.Init()
	go func() {
		t := time.Tick(time.Second)
		for range t {
			c := srv.Client()
			msg := c.NewMessage("kafka-topic", Message{ID: uuid.NewString(), Timestamp: time.Now().Unix()}, client.WithMessageContentType("application/json"))
			if err := c.Publish(context.TODO(), msg); err != nil {
				logger.Error(err)
			}
		}
	}()
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
