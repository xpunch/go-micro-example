package main

import (
	"context"

	_ "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"
	mopentracing "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	pb "github.com/xpunch/go-micro-example/v4/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	srv := micro.NewService(
		micro.Name("helloworld.cli"),
	)

	tracer, closer, err := jaegercfg.Configuration{ServiceName: "helloworld.cli",
		Reporter: &jaegercfg.ReporterConfig{LogSpans: true, CollectorEndpoint: "http://127.0.0.1:14268/api/traces"},
	}.NewTracer()
	if err != nil {
		logger.Fatal(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	srv.Init(micro.WrapClient(mopentracing.NewClientWrapper(tracer)))

	greeter := pb.NewHelloworldService("helloworld.srv", srv.Client())
	resp, err := greeter.Call(context.TODO(), &pb.Request{Name: "Client"})
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(resp)
}
