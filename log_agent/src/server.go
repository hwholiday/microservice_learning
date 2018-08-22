package src

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-micro/broker"
	"time"
	"strings"
	"log"

)
type LogServer struct{}

func Run(etcdAddr, nsqAddr, name, topic string) {
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = strings.Split(etcdAddr, ",")
	})
	nsqBroker := nsq.NewBroker(func(options *broker.Options) {
		options.Addrs = strings.Split(nsqAddr, ",")
	})
	server := micro.NewService(
		micro.Name(name),
		micro.Registry(registry),
		micro.Broker(nsqBroker),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	server.Init()
	// 订阅消息
	micro.RegisterSubscriber(topic, server.Server(), new(LogServer))

	if err := server.Run(); err != nil {
		log.Panic(err.Error())
	}
}

