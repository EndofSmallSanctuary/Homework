package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		os.Exit(9)
	}
}

func main() {
	if checkOpenPulls(os.Getenv("targetRepo")) {
		fmt.Println("Given repository has open Pull Requests which are waiting for merge")
		showWhatsHiding()
	} else {
		fmt.Println("Given repository has no open Pull Requests")
	}
}
