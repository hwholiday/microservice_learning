package main

import (
	"flag"
	"microservice_learning/db_agent/src"
)

func main() {
	flag.Parse()
	server := src.NewDbServer("1")
	server.Run()
}
