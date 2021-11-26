package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	_ "github.com/asim/go-micro/plugins/broker/kafka/v4"
	_ "github.com/asim/go-micro/plugins/broker/mqtt/v4"
	"github.com/google/uuid"
	pb "github.com/xpunch/go-micro-example/v4/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
)

type message struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	srv := micro.NewService(
		micro.Name("broker-client"),
	)
	srv.Init()
	var count int32
	go func() {
		t := time.NewTicker(5 * time.Second)
		for range t.C {
			m := &message{
				ID:        uuid.New().String(),
				Message:   fmt.Sprintf("No. %d", atomic.AddInt32(&count, 1)),
				Timestamp: time.Now().Unix(),
			}
			err := srv.Client().Publish(context.TODO(), client.NewMessage("micro.topic.json", m, client.WithMessageContentType("application/json")))
			if err != nil {
				logger.Error(err)
			}
		}
	}()
	go func() {
		event := micro.NewEvent("micro.topic.protobuf", srv.Client())
		t := time.NewTicker(5 * time.Second)
		for range t.C {
			m := &pb.Message{
				Id:        uuid.New().String(),
				Message:   fmt.Sprintf("No. %d", atomic.AddInt32(&count, 1)),
				Timestamp: time.Now().Unix(),
			}
			err := event.Publish(context.TODO(), m)
			if err != nil {
				logger.Error(err)
			}
		}
	}()
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
