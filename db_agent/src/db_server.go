package src

import (
	"time"
	"os"
	"os/signal"
	"syscall"
	"google.golang.org/grpc"
	"microservice_learning/protobuf/dbagent"
	"net"
	"fmt"
	utils "microservice_learning/common/server"

	"log"
)

type DbServer struct {

}

func NewDbServer(info string) *DbServer {
	return &DbServer{}
}

func (d *DbServer) Run(serv, reg string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
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
	s.Serve(lis)
}
