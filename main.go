package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// loadgin variables from .env file in YAML format

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file, error message ", err)
	}

}
