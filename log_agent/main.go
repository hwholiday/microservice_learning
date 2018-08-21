package main

import (
	"github.com/micro/go-micro"
	"log"
	"fmt"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-plugins/registry/etcdv3"
	"context"
	"microservice_learning/protobuf/logagent"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/broker"
	"time"
)

const topic = "server.log.data"

func main() {
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"http://127.0.0.1:2379"}
	})
	nsqBroker := nsq.NewBroker(func(options *broker.Options) {
		options.Addrs=[]string{"0.0.0.0:4152"}
	})
	server := micro.NewService(
		micro.Name("go.micro.srv.log"),
		micro.Registry(registry),
		micro.Broker(nsqBroker),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	server.Init()
	// 订阅消息
	micro.RegisterSubscriber(topic,server.Server(),ProcessEvent)

	if err := server.Run(); err != nil {
		log.Fatalf("srv run error: %v\n", err)
	}
}
func ProcessEvent(ctx context.Context, log *logagent.Log) error {
	fmt.Printf("Got event %+v\n", log)
	return nil
}

