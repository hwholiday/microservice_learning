package main

import "microservice_learning/db_agent/src"

func main() {
	dbServer := src.NewDbServer()
	dbServer.Run()
}
