package main

import (
	"log"
	"os"
	"path/filepath"
)

func checkCreateDir() {
	newpath := filepath.Join("Christmas Tree")
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
}

func createUserHTML(filename string, content []byte) {
	checkCreateDir()

	// f, err := os.Create("/tmp/dat2")
	// if
	err := os.WriteFile("Christmas Tree/"+filename+".html", content, 0644)
	if err != nil {
		log.Panic(err)
	}
}
