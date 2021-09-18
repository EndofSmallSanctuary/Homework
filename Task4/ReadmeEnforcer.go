package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// func decodeReadme(gitContent string) {
// 	if gitContent != "" {
// 		b64str, err := b64.StdEncoding.DecodeString(gitContent)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if len(b64str) > 5 {
// 			mastertail = b64str[(len(b64str)-1)-5 : len(b64str)-1]
// 		} else {
// 			mastertail = b64str
// 		}
// 	}
// }

func checkTaskValidity(gitContent string) string {
	data := url.Values{}
	data.Set("masterkey", os.Getenv("TOKEN"))
	data.Set("mastertail", string(gitContent))
	req, err := http.NewRequest("POST", os.Getenv("Raven"), strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
