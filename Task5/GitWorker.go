package main

import (
	"fmt"
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

	fmt.Println(requestSignature)
}

// func checkOpenPulls() int {

// }
