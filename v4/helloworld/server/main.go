package main

import (
	"context"
	"time"

	_ "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"
	pb "github.com/xpunch/go-micro-example/v4/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type HelloWorld struct{}

func (h *HelloWorld) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Message = "Hello " + req.Name
	// return errors.New("fake error")
	return nil
}

func main() {
	name := "helloworld.srv"
	srv := micro.NewService()
	tp, err := newTracerProvider(name, srv.Server().Options().Id, "http://jaeger-collector:14268/api/traces")
	if err != nil {
		logger.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	opts := []micro.Option{
		micro.Name(name),
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			logger.Fatal(err)
		}
	}()
	opts = append(opts, micro.WrapHandler(opentelemetry.NewHandlerWrapper()))
	srv.Init(opts...)
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(HelloWorld)); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

// newTracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func newTracerProvider(serviceName, serviceID, url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String("latest"),
			semconv.ServiceInstanceIDKey.String(serviceID),
		)),
	)
	return tp, nil
}
