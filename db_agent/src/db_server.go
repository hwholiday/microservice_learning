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
	"strings"
	"microservice_learning/common"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

type DbServer struct {
	pub      micro.Publisher
	db       *xorm.Engine
	sqlKey   string
	etcdAddr string
	nsqAddr  string
	name     string
	topic    string
}

func NewDbServer(etcd, nsq, name, topic, sql string) *DbServer {
	return &DbServer{etcdAddr: etcd, nsqAddr: nsq, name: name, topic: topic, sqlKey: sql}
}

func (d *DbServer) InitDb() {
	cli := common.InitEtcd(strings.Split(d.etcdAddr, ","))
	defer cli.Cli.Close()
	resp, err := cli.Get(d.sqlKey)
	if err != nil {
		log.Printf("read etcd %s fail", d.sqlKey)
		os.Exit(1)
	}
	if len(resp.Kvs) <= 0 {
		log.Printf("no %s info", d.sqlKey)
		os.Exit(1)
	}
	d.db, err = xorm.NewEngine("mysql", string(resp.Kvs[0].Value))
	if err != nil {
		log.Printf("init orm %s fail: %s", string(resp.Kvs[0].Value), err.Error())
		os.Exit(1)
	}
	err = d.db.Ping()
	if err != nil {
		log.Printf("init orm %s fail: %s", string(resp.Kvs[0].Value), err.Error())
		os.Exit(1)
	}
	d.db.ShowSQL(false)
}

func (d *DbServer) Run() {
	d.InitDb()
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = strings.Split(d.etcdAddr, ",")
	})
	nsqBroker := nsq.NewBroker(func(options *broker.Options) {
		options.Addrs = strings.Split(d.nsqAddr, ",")
	})
	service := micro.NewService(
		micro.Name(d.name),
		micro.Registry(registry),
		micro.Version("v1"),
		micro.Broker(nsqBroker),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	server.Init()
	d.pub = micro.NewPublisher(d.topic, service.Client())
	// Register handler
	//d.pub.Publish(context.TODO(), &logagent.Log{Time: time.Now().Unix(), Error: "errerre  ", Data: "db_agent启动成功", Filename: "main", Line: "35", Method: "main"}))
	dbagent.RegisterDbAgentServerHandler(service.Server(), d)

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}

}
