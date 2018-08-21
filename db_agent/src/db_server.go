package src

import (
	"fmt"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"microservice_learning/protobuf/dbagent"
)

type DbServer struct {

}

func NewDbServer(info string) *DbServer {
	return &DbServer{}
}

func (d *DbServer) Run() {
	registry := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Name("dbagent"),
		micro.Registry(registry),
	)
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	dbagent.RegisterDbAgentServerHandler(service.Server(), d)
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}



/*lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
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
	s.Serve(lis)*/