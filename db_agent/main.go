package main

import (
	"flag"
	"net"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"google.golang.org/grpc"
	utils "microservice_learning/common/server"
	"log"
	"microservice_learning/protobuf"
	"context"
)
var (
	serv = flag.String("service", "hello_service", "service name")
	port = flag.Int("port", 50001, "listening port")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)


func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}

	err = utils.Register(*serv, "127.0.0.1", *port, *reg, time.Second*10, 15)
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

	log.Printf("starting hello service at %d", *port)
	s := grpc.NewServer()
	protobuf.RegisterDbAgentServerServer(s, &DbServer{})
	s.Serve(lis)
}

type DbServer struct{
}
func (s *DbServer) GetOneTestUser(ctx context.Context, in *protobuf.StringValue) (*protobuf.TestUser, error) {
	fmt.Println(in.Value)
	return &protobuf.TestUser{Id:1,Account:"2",Password:"3"}, nil
}

func (s *DbServer) GetAllTestUser(ctx context.Context, in *protobuf.StringValue) (*protobuf.ListUser, error) {
	return nil, nil
}
