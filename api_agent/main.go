package main

import (
	"flag"
	"microservice_learning/api_agent/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
	"os"
	"microservice_learning/api_agent/router"
	"log"
)

var (
	name  = flag.String("name", "go.micro.srv.dbagent", "name")
	etcd  = flag.String("etcd", "http://127.0.0.1:2379", "register etcd address")
	nsq   = flag.String("nsq", "0.0.0.0:4152", "register nsq address")
	topic = flag.String("topic", "server.log.data", "register topic")
	addr  = flag.String("addr", "127.0.0.1:8089", "server address")
)

func init() {
	flag.Parse()
	server.InitServer(*etcd, *nsq, *name, *topic)
}
func main() {

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	router.SetRouters(g)
	s := &http.Server{
		Handler:        g,
		Addr:           *addr,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("server run :", *addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Panic(err)
		os.Exit(0)
	}
}
