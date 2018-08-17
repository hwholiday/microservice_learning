package main

import (
	"flag"
	"fmt"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	utils "microservice_learning/common/server"
	"microservice_learning/protobuf"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()
	fmt.Println("serv", *serv)
	r := utils.NewResolver(*serv)
	b := grpc.RoundRobin(r)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))
	if err != nil {
		panic(err)
	}
	fmt.Println("conn...")

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		client := protobuf.NewDbAgentServerClient(conn)
		resp, err := client.GetOneTestUser(context.Background(),&protobuf.StringValue{Value:"1"})
		if err!=nil{
			fmt.Println(err)
		}else {
			fmt.Println(resp)
		}
	}
}