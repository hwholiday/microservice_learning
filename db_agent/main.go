package main

import (
	"flag"
	"microservice_learning/db_agent/src"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	port = flag.Int("port", 50001, "listening port")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()
	server := src.NewDbServer("1")
	server.Run(*serv, *reg, *port)
}
