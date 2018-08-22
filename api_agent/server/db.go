package server

import (
	"microservice_learning/protobuf/dbagent"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-micro/broker"
	"time"
	"strings"
)


var Client *Server


type Server struct {
	DbAgent   dbagent.DbAgentServerService
	Pub       micro.Publisher
}


func InitServer(etcdAddr, nsqAddr,name,topic string) {
	Client=new(Server)
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
	Client.DbAgent= dbagent.NewDbAgentServerService(name, server.Client())
	Client.Pub = micro.NewPublisher(topic, server.Client())
	//s.Pub.Publish(context.Background(), &logagent.Log{Time: time.Now().Unix(), Error: "apiapiapi  ", Data: "api_agent启动成功", Filename: "apiapiapi", Line: "35", Method: "apiapiapi"})
	//ticker := time.NewTicker(1 * time.Second)
	//for range ticker.C {
	//	go func() {
	//		rsp, err := s.DbAgent.GetOneTestUser(context.Background(), &dbagent.StringValue{Value: "1"})
	//		if err != nil {
	//			fmt.Println(err.Error())
	//		} else {
	//			fmt.Println(rsp.String())
	//		}
	//	}()
	//}
}

