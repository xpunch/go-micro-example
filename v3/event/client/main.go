package main

import (
	"time"

	mgrpc "github.com/asim/go-micro/plugins/client/grpc/v3"
	mhttp "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
	pb "github.com/xpunch/go-micro-example/v3/event/proto"
	pbh "github.com/xpunch/go-micro-example/v3/helloworld/proto"
)

func main() {
	srv := micro.NewService(
		micro.Server(mhttp.NewServer()),
		micro.Client(mgrpc.NewClient()),
		micro.Name("web"),
		micro.Address(":80"),
	)
	srv.Init()
	accessEvent := micro.NewEvent("accesslogs", srv.Client())
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger(), AccessLogMiddleware(accessEvent))
	helloworldService := pbh.NewHelloworldService("helloworld", srv.Client())
	statisticsService := pb.NewStatisticsService("statistics", srv.Client())
	router.POST("/helloworld", func(ctx *gin.Context) {
		resp, err := helloworldService.Call(ctx, &pbh.Request{Name: ctx.Query("user")})
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		ctx.JSON(200, resp)
	})
	router.GET("/statistics", func(ctx *gin.Context) {
		method := ctx.Query("method")
		resp, err := statisticsService.Statistics(ctx, &pb.StatisticsRequest{Method: &method})
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		ctx.JSON(200, resp)
	})
	if err := micro.RegisterHandler(srv.Server(), router); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Error(err)
	}
}

func AccessLogMiddleware(event micro.Event) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.EscapedPath()
		ctx.Next()
		method := ctx.Request.Method
		latency := time.Since(start)
		err := event.Publish(ctx, &pb.AccessEvent{
			Status:    uint32(ctx.Writer.Status()),
			Method:    method,
			Path:      path,
			Ip:        ctx.ClientIP(),
			Latency:   int64(latency),
			Timestamp: start.Unix(),
		})
		if err != nil {
			logger.Warn(err)
		}
	}
}
