package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type repoContent struct {
	Taskname string `json:"path"`
	DeepLink string `json:"html_url"`
	Type     string `json:"type"`
}

func prepareRepoLink() string {
	val := os.Getenv("Repo")
	return val
}

func obtainTaskList() {

	repotail := strings.Split(os.Getenv("Repo"), "github.com")
	if len(repotail) < 2 {
		log.Panic("incorrect repository link")
		return
	}

	resp, err := http.Get("https://api.github.com/repos" + repotail[1] + "/contents/")
	fmt.Println(os.Getenv("Repo") + "/contents/")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	data := []repoContent{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i < len(data); i++ {
		fmt.Println(data[i])
	}
}
