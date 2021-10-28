package main

import "github.com/micro/go-micro"

func main() {
	srv := micro.NewService()
	srv.Init()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
