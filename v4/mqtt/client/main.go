package main

import (
	"context"
	"fmt"
	"time"

	"github.com/asim/go-micro/plugins/broker/mqtt/v4"
	"github.com/google/uuid"
	pb "github.com/xpunch/go-micro-example/v4/mqtt/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
)

type message struct {
	ID        string `json:"id"`
	Message   string `jsno:"msg"`
	Timestamp int64  `json:"ts"`
}

func main() {
	srv := micro.NewService(
		micro.Name("mqtt-client"),
		micro.Broker(mqtt.NewBroker()),
	)
	srv.Init()
	go func() {
		t, count := time.NewTicker(5*time.Second), 1
		for range t.C {
			m := &message{
				ID:        uuid.New().String(),
				Message:   fmt.Sprintf("No. %d", count),
				Timestamp: time.Now().Unix(),
			}
			err := srv.Client().Publish(context.TODO(), client.NewMessage("micro.broker.mqtt.json", m, client.WithMessageContentType("application/json")))
			if err != nil {
				logger.Error(err)
			}
			count++
		}
	}()
	go func() {
		event := micro.NewEvent("micro.broker.mqtt.protobuf", srv.Client())
		t, count := time.NewTicker(5*time.Second), 1
		for range t.C {
			m := &pb.Message{
				Id:        uuid.New().String(),
				Message:   fmt.Sprintf("No. %d", count),
				Timestamp: time.Now().Unix(),
			}
			err := event.Publish(context.TODO(), m)
			if err != nil {
				logger.Error(err)
			}
			count++
		}
	}()
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
