package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

type Gitpulls struct {
	Url     int64       `json:"id"`
	GitUser GitUsers    `json:"user"`
	Labels  []GitLabels `json:"labels"`
}

type GitUsers struct {
	Id     int64  `json:"id"`
	Login  string `json:"login"`
	Avatar string `json:"avatar_url"`
}

type GitLabels struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OutputNode struct {
	User            GitUsers    `json:"user"`
	Lables          []GitLabels `json:"labels"`
	TotalPrs        int32       `json:"totalprs"`
	TotalPrsWLabels int32
}

var sessionPulls = []Gitpulls{}
var outputNodes = []OutputNode{}

func showWhatsHiding() {
	sort.Slice(sessionPulls, func(i, j int) bool {
		return sessionPulls[i].GitUser.Id < sessionPulls[j].GitUser.Id
	})
	fmt.Println("Analizing open PRs statistic...")
	currentId := sessionPulls[0].GitUser.Id
	outputNodes = append(outputNodes, OutputNode{User: sessionPulls[0].GitUser, TotalPrs: 0, TotalPrsWLabels: 0})

	for _, pull := range sessionPulls {
		if pull.GitUser.Id == currentId {
			outputNodes[len(outputNodes)-1].TotalPrs++
			if len(pull.Labels) > 0 {
				outputNodes[len(outputNodes)-1].Lables = append(outputNodes[len(outputNodes)-1].Lables, pull.Labels...)
				outputNodes[len(outputNodes)-1].TotalPrsWLabels++
			}
		} else {
			currentId = pull.GitUser.Id
			outputNodes = append(outputNodes, OutputNode{User: pull.GitUser, TotalPrs: 1})
		}
	}

	fmt.Println("Most productive contributors: ")
	for _, node := range outputNodes {
		fmt.Println(node.User.Login + " made " + fmt.Sprintf("%d", node.TotalPrs) + " PRs, of which " + fmt.Sprintf("%d", node.TotalPrsWLabels) + " were made with labels")
		createUserHTML(node.User.Login, testRequest(node))
	}

	fmt.Println("")
	fmt.Println("I have a strong feeling, that you should check out your Christmas Tree :)")

}

func testRequest(pull OutputNode) []byte {

	payload, err := json.Marshal(pull)
	if err != nil {
		log.Panic(err)
	}

	req, err := http.NewRequest("POST", os.Getenv("Raven"), bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Cannot establish connection with Raven server")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	result := decode64(string(body))
	return result
}

func decode64(base64 string) []byte {
	result, err := b64.StdEncoding.DecodeString(base64)
	if err != nil {
		log.Panic(err)
	}
	return result
}
