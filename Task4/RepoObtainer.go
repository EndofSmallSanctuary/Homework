package main

import (
	b64 "encoding/base64"
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

func obtainTaskList() {

	fmt.Println(prepareRequestSignature(""))

	resp, err := http.Get(apisignature)
	fmt.Println(os.Getenv("Repo") + "/contents/")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	sessionTasks = []repoContent{}
	err = json.Unmarshal([]byte(body), &sessionTasks)
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
		req, err := http.Get(prepareRequestSignature(sessionTasks[i].Taskname + "/README.md"))
		if err != nil {
			log.Fatal(err)
		}

		defer req.Body.Close()

		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			log.Fatal(err)
		}

		reqData := make(map[string]interface{})
		err = json.Unmarshal([]byte(body), &reqData)
		if err != nil {
			log.Fatal(err)
		}

		if reqData["content"] != nil {
			basedContent := reqData["content"].(string)

			if basedContent != "" {
				b64str, err := b64.StdEncoding.DecodeString(basedContent)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(b64str))
			}
		}
	}
}
