package main

import (
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
	pb "github.com/xpunch/go-micro-example/v3/helloworld/proto"
)

func main() {
	srv := micro.NewService(
		micro.Server(http.NewServer()),
		micro.Client(grpc.NewClient()),
		micro.Name("web"),
		micro.Address(":80"),
	)
	srv.Init()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	helloworldService := pb.NewHelloworldService("helloworld", srv.Client())
	router.POST("/helloworld", func(ctx *gin.Context) {
		var req struct {
			User string `json:"user"`
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(400, err)
			return
		}
		resp, err := helloworldService.Call(ctx, &pb.Request{Name: req.User})
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
		logger.Fatal(err)
	}
}
