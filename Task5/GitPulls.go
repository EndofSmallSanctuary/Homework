package main

import (
	"fmt"
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
	User     GitUsers
	Lables   []GitLabels
	TotalPrs int32
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

	for _, pull := range outputNodes {
		fmt.Println(pull)
	}

}
