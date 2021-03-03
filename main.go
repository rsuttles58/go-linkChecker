package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// inputProcessor asks the user what url they want to check and returns that value as a string
func inputProcessor() string {
	fmt.Println("What link would you like to check?")
	fmt.Println("Be sure to include 'https://' before the url")
	var url string
	fmt.Scanln(&url)
	return url
}

// urlChecker receives a string and returns true or false if the string includes "https://"
func urlChecker(url string) {
	if strings.Contains(url, "https://") == false {
		log.Fatalln("invalid url string provided.")
	}
}

// FetchURL will take a passed url string and send an HTTP GET request to that url
func FetchURL(url string) (*http.Response, error) {
	return http.Get(url)
}

func main() {
	url := inputProcessor()
	urlChecker(url)

	resp, err := FetchURL(url)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.StatusCode)
}
