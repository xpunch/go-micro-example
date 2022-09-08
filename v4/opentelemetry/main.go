package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-micro/plugins/v4/registry/etcd"
	_ "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"

	pb "github.com/xpunch/go-micro-example/v4/proto"
)

const (
	id      = 1
	name    = "opentelemetry.srv"
	version = "latest"
)

func main() {
	srv := micro.NewService(
		micro.Name(name),
	)
	tp, err := tracerProvider("http://47.103.38.143:14268/api/traces")
	// tp, err := tracerProvider("http://jaeger-collector.default.svc.cluster.local:14268/api/traces")
	if err != nil {
		logger.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	srv.Init(micro.WrapClient(opentelemetry.NewClientWrapper()))
	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			logger.Fatal(err)
		}
	}(context.TODO())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		tctx, span := tp.Tracer("/").Start(ctx, "request")
		defer span.End()
		fmt.Fprintln(w, "hello world")

		md := make(metadata.Metadata, 10)
		otel.GetTextMapPropagator().Inject(tctx, &metadataSupplier{metadata: &md})
		fmt.Println(md)
		ctx = metadata.NewContext(tctx, md)
		greeter := pb.NewHelloworldService("helloworld.srv", srv.Client())
		resp, err := greeter.Call(ctx, &pb.Request{Name: "Client"})
		if err != nil {
			logger.Error(err)
		}
		logger.Info(resp)
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
			semconv.ServiceVersionKey.String(version),
		)),
	)
	return tp, nil
}
