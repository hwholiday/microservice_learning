package main

import (
	"flag"
	"github.com/micro/go-micro"
	"fmt"
	"microservice_learning/protobuf/dbagent"
	"golang.org/x/net/context"
	"time"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")

)

func main() {
	registry := etcdv3.NewRegistry()
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("dbagent"),micro.Registry(registry))

	// Create new greeter client
	greeter := dbagent.NewDbAgentServerService("dbagent", service.Client())

	// Call the greeter
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		rsp, err := greeter.GetOneTestUser(context.TODO(), &dbagent.StringValue{Value:"1"})
		if err != nil {
			fmt.Println(err)
		}
		// Print response
		fmt.Println(rsp.String())
	}

	/*flag.Parse()
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
	}*/
}