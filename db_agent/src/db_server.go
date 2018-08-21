package src

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"microservice_learning/protobuf/dbagent"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/broker/nsq"
	"microservice_learning/protobuf/logagent"
	"time"
	"context"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/broker"
)

type DbServer struct {
}

func NewDbServer()*DbServer  {
	return &DbServer{}
}
func (d *DbServer)Run() {
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"http://127.0.0.1:2379"}
	})
	nsqBroker := nsq.NewBroker(func(options *broker.Options) {
		options.Addrs=[]string{"127.0.0.1:4152"}
	})
	service := micro.NewService(
		micro.Name("service.howie"),
		micro.Registry(registry),
		micro.Broker(nsqBroker),
	)
	server.Init()
	sub:= micro.NewPublisher("server.log.data", service.Client())
	// Register handler
	sub.Publish(context.TODO(), &logagent.Log{Time:time.Now().Unix(),Error:"errerre  ",Data:"db_agent启动成功",Filename:"main",Line:"35",Method:"main"})
	dbagent.RegisterDbAgentServerHandler(service.Server(), d)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

/*lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}
	err = utils.Register(serv, "127.0.0.1", port, reg, time.Second*10, 15)
	if err != nil {
		panic(err)
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		utils.UnRegister()
		os.Exit(1)
	}()
	log.Printf("starting hello service at %d", port)
	s := grpc.NewServer()
	dbagent.RegisterDbAgentServerServer(s, d)
	s.Serve(lis)*/
