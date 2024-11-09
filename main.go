package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

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

	connStr := godotenv.Load("eventHubConnString")

	hub, err := eventhub.NewHubFromConnectionString(connStr, err)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// send test message

	err = hub.Send(ctx, eventhub.NewEventFromString("hello, world!"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
