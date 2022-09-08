package main

import (
	"context"
	"sync"

	_ "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"
	mhystrix "github.com/go-micro/plugins/v4/wrapper/breaker/hystrix"
	pb "github.com/xpunch/go-micro-example/v4/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	srv := micro.NewService(
		micro.Name("helloworld.cli"),
	)
	opts := []micro.Option{}
	mhystrix.ConfigureDefault(mhystrix.CommandConfig{MaxConcurrentRequests: 1})
	opts = append(opts, micro.WrapClient(mhystrix.NewClientWrapper(mhystrix.WithFilter(func(c context.Context, e error) bool {
		return e == nil || e.Error() == "ignore me"
	}))))
	srv.Init(opts...)
	greeter := pb.NewHelloworldService("helloworld.srv", srv.Client())
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := greeter.Call(context.TODO(), &pb.Request{Name: "Client"})
			if err != nil {
				logger.Error(err)
			} else {
				logger.Info(resp)
			}
		}()
	}
	wg.Wait()
}
