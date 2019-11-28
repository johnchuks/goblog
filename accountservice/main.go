package main

import (
	"github.com/johnchuks/goblog/accountservice/dbclient"
	"fmt"
	"github.com/johnchuks/goblog/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	intializeBoltClient()
	service.StartWebServer("3200")
}

func intializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
