package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"microservice_learning/common"
)

func main() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 典型读取操作，默认分区可以使用空字符串表示
	cli := common.InitEtcd([]string{"http://127.0.0.1:2379"})
	sql := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", cfg.Section("db_agent_test").Key("user").String(), cfg.Section("db_agent_test").Key("passwod").String(), cfg.Section("db_agent_test").Key("url").String(), cfg.Section("db_agent_test").Key("prot").String(), cfg.Section("db_agent_test").Key("database").String())
	_, err = cli.Put("db_agent_test", sql)
	if err != nil {
		fmt.Printf("Fail to etcd db_agent_test %s\n", sql)
		os.Exit(1)
	}
	fmt.Printf("success to etcd db_agent_test %s\n", sql)
	cli.Cli.Close()
}
