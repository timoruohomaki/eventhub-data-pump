package main

import (
	"fmt"
	"github.com/Azure/azure-event-hubs-go/v3"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {

	// loading variables from .env file in YAML format
	// .env added in gitignore

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file, error message ", err)
	}

	build := 3
	thisHost, _ := os.Hostname()
	buildEnv := os.Getenv("buildEnv")

	fmt.Println()
	fmt.Println("=============================================")
	fmt.Println("=  Starting EventHub Data Pump...           =")
	fmt.Println("=============================================")
	fmt.Println("  Build version:    ", strconv.Itoa(build))
	fmt.Println("  Host name:        ", thisHost)
	fmt.Println("  Environment:      ", buildEnv)
	fmt.Println("=============================================")
	fmt.Println()

	// Connecting Azure Event Hub

	connStr := godotenv.Load("eventHubConnString")

	hub, err := eventhub.NewHubFromConnectionString(connStr, err)

}
