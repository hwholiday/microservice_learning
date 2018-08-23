package main

import (
	"flag"
	"microservice_learning/api_agent/server"
	"github.com/gin-gonic/gin"
	"microservice_learning/api_agent/router"
)

var (
	name  = flag.String("name", "go.micro.srv.dbagent", "name")
	etcd  = flag.String("etcd", "http://127.0.0.1:2379", "register etcd address")
	nsq   = flag.String("nsq", "0.0.0.0:4152", "register nsq address")
	topic = flag.String("topic", "server.log.data", "register topic")
	addr  = flag.String("addr", "127.0.0.1:8089", "server address")
)


func main() {
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	router.SetRouters(g)
	server.InitServer(g,*etcd, *nsq, *name, *topic)
}
