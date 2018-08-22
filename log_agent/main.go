package main

import (
	"flag"
	"microservice_learning/log_agent/src"
)

var (
	name = flag.String("name", "go.micro.srv.log", "name")
	etcd  = flag.String("etcd", "http://127.0.0.1:2379", "register etcd address")
	nsq  = flag.String("nsq", "0.0.0.0:4152", "register nsq address")
	topic  = flag.String("topic", "server.log.data", "register topic")

)

func main() {
	flag.Parse()
	src.Run(*etcd,*nsq,*name,*topic)
}
