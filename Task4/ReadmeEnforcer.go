package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var mastertail []byte

func decodeReadme(gitContent string) {
	if gitContent != "" {
		b64str, err := b64.StdEncoding.DecodeString(gitContent)
		if err != nil {
			log.Fatal(err)
		}

		if len(b64str) > 5 {
			mastertail = b64str[(len(b64str)-1)-5 : len(b64str)-1]
		} else {
			mastertail = b64str
		}
	}
}

// func retrieveTaskStatus() bool {
// 	lastind := len(rcontent) - 1
// 	if lastind > 10 {
// 		substr := rcontent[lastind-10 : lastind]
// 		bytes := []byte(substr)
// 		fmt.Println(bytes)

// 	}
// 	return false
// }

func checkTaskValidity(gitContent string) {
	decodeReadme(gitContent)
	data := url.Values{}
	data.Set("masterkey", string(mastertail))
	req, err := http.NewRequest("POST", "http://localhost:8080/testKey", strings.NewReader(data.Encode()))
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(body)
}
