package common

import (
	"time"
	"golang.org/x/net/context"
	"github.com/coreos/etcd/clientv3"
	"log"
	"os"
)

type EtcdCli struct {
	Cli *clientv3.Client
}

func InitEtcd(addr []string) (*EtcdCli) {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Print(err)
		os.Exit(0)
	}
	return &EtcdCli{etcdCli}
}

func (e *EtcdCli)Get(key string) (resp *clientv3.GetResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err = e.Cli.Get(ctx, key)
	cancel()
	return
}

func (e *EtcdCli)Put(key string, value string) (resp *clientv3.PutResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err = e.Cli.Put(ctx, key, value)
	cancel()
	return
}
