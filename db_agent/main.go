package main

import (
	"microservice_learning/db_agent/src"
	"flag"
)

var (
	name  = flag.String("name", "go.micro.srv.dbagent", "name")
	etcd  = flag.String("etcd", "http://127.0.0.1:2379", "register etcd address")
	nsq   = flag.String("nsq", "0.0.0.0:4150", "register nsq address")
	topic = flag.String("topic", "server.log.data", "register topic")
)

func main() {

	flag.Parse()
	dbServer := src.NewDbServer(*etcd, *nsq, *name, *topic)
	dbServer.Run()
}
