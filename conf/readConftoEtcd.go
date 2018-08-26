package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

func main() {
	cfg, err := ini.Load("C:\\Project\\Go\\src\\microservice_learning\\conf\\conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())
}
