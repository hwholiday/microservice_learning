package server

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
	"strings"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-web"
	"time"
	"github.com/micro/go-micro/client"
	"microservice_learning/protobuf/dbagent"
	"microservice_learning/protobuf/logagent"
	"fmt"
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"os"
)


var Client *Server


type Server struct {
	DbAgent   dbagent.DbAgentServerService
	Pub       micro.Publisher
}


func InitServer(r *gin.Engine,args ...string) {
	if len(args)<4{
		os.Exit(1)
		return
	}
	Client=new(Server)
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = strings.Split(args[0], ",")
	})
	nsqBroker := nsq.NewBroker(func(options *broker.Options) {
		options.Addrs = strings.Split(args[1], ",")
	})
	service := web.NewService(
		web.Name("go.micro.api.api"),
		web.Registry(registry),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
	)
	service.Init()
	cli:=client.NewClient(
		client.Broker(nsqBroker),
		client.Registry(registry),
	)
	Client.DbAgent= dbagent.NewDbAgentServerService(args[2],cli)
	Client.Pub = micro.NewPublisher(args[3], cli)
	Client.Pub.Publish(context.Background(), &logagent.Log{Time: time.Now().Unix(), Error: "apiapiapi  ", Data: "api_agent启动成功", Filename: "apiapiapi", Line: "35", Method: "apiapiapi"})

	service.Handle("/", r)
	go func() {
		rsp, err := Client.DbAgent.GetOneTestUser(context.Background(), &dbagent.StringValue{Value: "1"})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(rsp.String())
		}
	}()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}


