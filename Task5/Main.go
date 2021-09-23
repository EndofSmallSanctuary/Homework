package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var repositoryInitial string = "http://github.com/"

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		os.Exit(9)
	}
}

func main() {

	var repoOwner string
	var repoChosen string

	fmt.Println("Please, let me know username, whose repository we are going to check")
	fmt.Scanln(&repoOwner)
	fmt.Println("Please, let me know which repository of " + repoOwner + " collection, we are going to check")
	fmt.Scanln(&repoChosen)

	repositoryInitial += repoOwner + "/" + repoChosen + "/"

	if checkOpenPulls(repositoryInitial) {
		fmt.Println("Given repository has open Pull Requests which are waiting for merge")
		fmt.Println("")
		showWhatsHiding()
	} else {
		fmt.Println("Given repository has no open Pull Requests")
	}
}
