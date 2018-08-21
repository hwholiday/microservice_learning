package main

import (
	"github.com/micro/go-micro"
	"log"
	"fmt"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-plugins/registry/etcdv3"
	"context"
	"microservice_learning/protobuf/logagent"
)

const topic = "server.log.data"

func main() {
	registry := etcdv3.NewRegistry()
	nsqBroker:=nsq.NewBroker()
	server := micro.NewService(
		micro.Name("go.micro.srv.server"),
		micro.Version("v1"),
		micro.Registry(registry),
		micro.Broker(nsqBroker),
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

