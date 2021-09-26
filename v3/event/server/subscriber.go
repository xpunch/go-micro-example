package main

import (
	"context"
	"sync/atomic"

	"github.com/asim/go-micro/v3/logger"
	pb "github.com/xpunch/go-micro-example/v3/event/proto"
)

type subscriber struct{}

func (s *subscriber) AccessLog(ctx context.Context, l *pb.AccessEvent) error {
	defer func() {
		if e := recover(); e != nil {
			logger.Error("panic: %v", e)
		}
	}()
	logger.Info(l)
	atomic.AddInt64(&AccessCount, 1)
	mu.Lock()
	defer mu.Unlock()
	v, ok := AccessMethods[l.Method]
	if ok {
		AccessMethods[l.Method] = v + 1
	} else {
		AccessMethods[l.Method] = 1
	}
	return nil
}
