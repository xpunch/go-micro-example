package main

import (
	"os"
	"os/signal"

	"github.com/asim/go-micro/plugins/config/source/etcd/v3"
	"github.com/asim/go-micro/v3/config"
	"github.com/asim/go-micro/v3/logger"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	etcdSource := etcd.NewSource(
		etcd.WithAddress("127.0.0.1:2379"),
		etcd.WithPrefix("/key"),
		etcd.StripPrefix(true),
	)
	err := config.Load(etcdSource)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(string(config.Bytes()))

	w, err := etcdSource.Watch()
	if err != nil {
		logger.Fatal(err)
	}
	defer w.Stop()
	go func() {
		for {
			cs, err := w.Next()
			if err != nil {
				logger.Error(err)
				break
			}
			logger.Infof("[watcher] %s", cs.Data)
		}
	}()
	<-exit
}
