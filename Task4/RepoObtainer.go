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

var apisignature string = ""
var sessionTasks []repoContent

type repoContent struct {
	Taskname string `json:"path"`
	DeepLink string `json:"html_url"`
	Type     string `json:"type"`
}

func prepareRepoLink() string {
	val := os.Getenv("Repo")
	return val
}

func prepareRequestSignature(postfix string) string {
	if apisignature == "" {
		repotail := strings.Split(prepareRepoLink(), "github.com")
		if len(repotail) < 2 {
			log.Panic("incorrect repository link")
			return apisignature
		} else {
			apisignature = "https://api.github.com/repos" + repotail[1] + "/contents/"
		}
	}
	fmt.Println(apisignature + postfix)
	return apisignature + postfix
}

func prepareAuthorizedRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("Authorization", os.Getenv("TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	return body
}

func obtainTaskList() {

	fmt.Println(prepareRequestSignature(""))

	body := prepareAuthorizedRequest(prepareRequestSignature(""))
	sessionTasks = []repoContent{}
	err := json.Unmarshal([]byte(body), &sessionTasks)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func parseTasks() {

	obtainTaskList()

	for i := 1; i < len(sessionTasks); i++ {

		fmt.Println(sessionTasks[i].Taskname)

		//Что это такое ??? Что он делает ахахахах
		body := prepareAuthorizedRequest(prepareRequestSignature(sessionTasks[i].Taskname + "/README.md"))
		reqData := make(map[string]interface{})
		err := json.Unmarshal([]byte(body), &reqData)
		if err != nil {
			log.Fatal(err)
		}

		if reqData["content"] != nil {
			checkTaskValidity(reqData["content"].(string))
			// decodeReadme(reqData["content"].(string))
			//retrieveTaskStatus()
		}
	}
}
