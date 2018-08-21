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
)

type DbServer struct {
}

func NewDbServer()*DbServer  {
	return &DbServer{}
}
func (d *DbServer)Run() {

	registry := etcdv3.NewRegistry()
	nsqBroker := nsq.NewBroker()
	service := micro.NewService(
		micro.Name("go.micro.srv.server"),
		micro.Version("v1"),
		micro.Registry(registry),
		micro.Broker(nsqBroker),
	)
	server.Init()
	sub:= micro.NewPublisher("server.log.data", service.Client())
	// Register handler
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			sub.Publish(context.TODO(), &logagent.Log{Time:time.Now().Unix(),Error:"errerre  ",Data:"1231231231",Filename:"main",Line:"35",Method:"main"})
		}
	}()
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
