package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Input format  = https://github.com/$user/$repo
var requestSignature = "https://api.github.com/repos/{user}/{repo}/pulls"
var targetRepo = ""

func prepareRequest() {
	if targetRepo != "" {
		repoSlice := strings.Split(targetRepo, "/")
		requestSignature = strings.Replace(requestSignature, "{user}", repoSlice[3], 1)
		requestSignature = strings.Replace(requestSignature, "{repo}", repoSlice[4], 1)
	}
}

func checkOpenPulls(input string) bool {
	targetRepo = input
	prepareRequest()

	req, err := http.Get(requestSignature)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panic(err)
	}

	sessionPulls = []Gitpulls{}
	err = json.Unmarshal([]byte(body), &sessionPulls)
	if err != nil {
		fmt.Println(err)
	}
	if len(sessionPulls) > 0 {
		return true
	} else {
		return false
	}

}
