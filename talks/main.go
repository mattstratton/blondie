package main

import (
	"fmt"

	"github.com/mattstratton/blondie/talks/dbclient"
	"github.com/mattstratton/blondie/talks/service"
)

var appName = "talkservice"

func main() {

	fmt.Printf("Starting %v\n", appName)
	initializeMongoClient()
	service.StartWebServer("8083")

}

// connect to db
func initializeMongoClient() {
	service.DBClient = &dbclient.MongoClient{}
	service.DBClient.OpenMongoDB()
	service.DBClient.Seed()
}
