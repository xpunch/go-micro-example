package main

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/asim/go-micro/plugins/client/grpc/v4"
	_ "github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

func main() {
	srv := micro.NewService(
		micro.Server(http.NewServer()),
		micro.Client(grpc.NewClient()),
		micro.Name("gateway"),
		micro.Address(":8080"),
	)
	srv.Init()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.POST("/:service/:endpoint", func(ctx *gin.Context) {
		service, endpoint := ctx.Param("service"), ctx.Param("endpoint")
		defer ctx.Request.Body.Close()
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error(err)
			ctx.AbortWithStatusJSON(500, err.Error())
			return
		}
		var request json.RawMessage
		if len(data) > 0 {
			d := json.NewDecoder(strings.NewReader(string(data)))
			d.UseNumber()
			if err := d.Decode(&request); err != nil {
				logger.Error(err)
				ctx.AbortWithStatusJSON(500, err.Error())
				return
			}
		}
		c := srv.Client()
		var response json.RawMessage
		if err := c.Call(ctx, c.NewRequest(service, endpoint, request, client.WithContentType("application/json")), &response); err != nil {
			logger.Error(err)
			ctx.AbortWithStatusJSON(500, err.Error())
			return
		}
		ctx.JSON(200, response)
	})
	router.GET("/:service/nodes", func(ctx *gin.Context) {
		services, err := srv.Options().Registry.GetService(ctx.Param("service"))
		if err != nil {
			logger.Error(err)
			ctx.AbortWithStatusJSON(500, err.Error())
			return
		}
		if len(services) == 0 {
			ctx.AbortWithStatusJSON(400, "service not found")
			return
		}
		nodes := make([]*registry.Node, 0)
		for _, s := range services {
			nodes = append(nodes, s.Nodes...)
		}
		ctx.JSON(200, nodes)
	})
	if err := micro.RegisterHandler(srv.Server(), router); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
