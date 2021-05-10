package main

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/cmd"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
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
		s, err := srv.Options().Registry.GetService(service)
		if err != nil {
			logger.Error(err)
			ctx.AbortWithStatusJSON(400, err.Error())
			return
		}
		if len(s) == 0 {
			ctx.AbortWithStatusJSON(400, "service not found")
			return
		}
		defer ctx.Request.Body.Close()
		form, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error(err)
			ctx.AbortWithStatusJSON(400, err.Error())
			return
		}
		var request json.RawMessage
		d := json.NewDecoder(strings.NewReader(string(form)))
		d.UseNumber()
		if err := d.Decode(&request); err != nil {
			logger.Error(err)
			ctx.AbortWithStatusJSON(400, err.Error())
			return
		}
		c := *cmd.DefaultOptions().Client
		var response json.RawMessage
		if err := c.Call(ctx, c.NewRequest(service, endpoint, request, client.WithContentType("application/json")), &response); err != nil {
			logger.Error(err)
			ctx.AbortWithStatusJSON(400, err.Error())
			return
		}
		ctx.JSON(200, response)
	})
	if err := micro.RegisterHandler(srv.Server(), router); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
