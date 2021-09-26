package main

import (
	"context"
	"sync/atomic"

	pb "github.com/xpunch/go-micro-example/v3/event/proto"
)

type handler struct{}

func (h *handler) Statistics(ctx context.Context, in *pb.StatisticsRequest, out *pb.StatisticsReply) error {
	if in.Method == nil || len(*in.Method) == 0 {
		out.AccessCount = atomic.LoadInt64(&AccessCount)
	} else {
		mu.RLock()
		defer mu.RUnlock()
		v, ok := AccessMethods[*in.Method]
		if ok {
			out.AccessCount = v
		}
	}
	return nil
}
