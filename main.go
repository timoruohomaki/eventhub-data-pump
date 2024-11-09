package main

import (
	"fmt"
	"log"
	"os"

	"strconv"

	// "github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/joho/godotenv"
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

	// type connStr string

	connStr = connStr(os.Getenv("CONNSTR"))

	fmt.Println(os.Getenv("CONNSTR"))

	// evProducer, err := azeventhubs.NewProducerClientFromConnectionString(connStr, "", nil)

}
