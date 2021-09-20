package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

type Gitpulls struct {
	Url     int32       `json:"id"`
	GitUser GitUsers    `json:"user"`
	Labels  []GitLabels `json:"labels"`
}

type GitUsers struct {
	Id    int32  `json:"id"`
	Login string `json:"login"`
}

type GitLabels struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OutputNode struct {
	User     GitUsers    `json:"user"`
	Lables   []GitLabels `json:"labels"`
	TotalPrs int32       `json:"totalprs"`
}

var sessionPulls = []Gitpulls{}
var outputNodes = []OutputNode{}

func showWhatsHiding() {
	sort.Slice(sessionPulls, func(i, j int) bool {
		return sessionPulls[i].GitUser.Id < sessionPulls[j].GitUser.Id
	})
	fmt.Println("Analizing open PRs statistic...")
	currentId := sessionPulls[0].GitUser.Id
	outputNodes = append(outputNodes, OutputNode{User: sessionPulls[0].GitUser, TotalPrs: 0})

	for _, pull := range sessionPulls {
		if pull.GitUser.Id == currentId {
			outputNodes[len(outputNodes)-1].TotalPrs++
			if len(pull.Labels) > 0 {
				outputNodes[len(outputNodes)-1].Lables = append(outputNodes[len(outputNodes)-1].Lables, pull.Labels...)
			}
		} else {
			currentId = pull.GitUser.Id
			outputNodes = append(outputNodes, OutputNode{User: pull.GitUser, TotalPrs: 1})
		}
	}

	fmt.Println(len(outputNodes))

	testRequest(outputNodes)

}

func testRequest(pull []OutputNode) string {

	payload, err := json.Marshal(outputNodes)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(payload))

	req, err := http.NewRequest("POST", os.Getenv("Raven"), bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Cannot establish connection with Raven server"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	result := string(body)
	return result
}
