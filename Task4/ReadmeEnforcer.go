package main

import (
	b64 "encoding/base64"
	"fmt"
	"log"
)

var mastertail []byte

func decodeReadme(gitContent string) {
	if gitContent != "" {
		b64str, err := b64.StdEncoding.DecodeString(gitContent)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(b64str)
		if len(b64str) > 5 {
			mastertail = b64str[(len(b64str)-1)-5 : len(b64str)-1]
		} else {
			fmt.Println(len(mastertail))
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

// func checkTaskValidity(){
// 	checkStr =
// }
