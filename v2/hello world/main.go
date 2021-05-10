package main

import (
	micro "github.com/micro/go-micro/v2"
)

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
	)

	// initialise flags
	service.Init()

	// start the service
	service.Run()
}
