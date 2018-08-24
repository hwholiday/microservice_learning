package src

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"microservice_learning/protobuf/dbagent"
	"github.com/micro/go-micro/server"
	"time"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-micro/broker"
	"log"
	"github.com/go-xorm/xorm"
	"fmt"
	"microservice_learning/protobuf/logagent"
	"context"
	"strings"
)

type DbServer struct {
	pub      micro.Publisher
	db       *xorm.Engine
	etcdAddr string
	nsqAddr  string
	name     string
	topic    string
}

func NewDbServer(etcd,nsq,name,topic string) *DbServer {
	return &DbServer{etcdAddr:etcd,nsqAddr:nsq,name:name,topic:topic}
}
func (d *DbServer) Run() {
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = strings.Split(d.etcdAddr, ",")
	})
	nsqBroker := nsq.NewBroker(func(options *broker.Options) {
		options.Addrs = strings.Split(d.nsqAddr, ",")
	})
	service := micro.NewService(
		micro.Name(d.name),
		micro.Registry(registry),
		micro.Broker(nsqBroker),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	server.Init()
	d.pub = micro.NewPublisher(d.topic, service.Client())
	// Register handler
	fmt.Println(d.pub.Publish(context.TODO(), &logagent.Log{Time: time.Now().Unix(), Error: "errerre  ", Data: "db_agent启动成功", Filename: "main", Line: "35", Method: "main"}))
	dbagent.RegisterDbAgentServerHandler(service.Server(), d)

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
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
