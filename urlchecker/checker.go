package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {

	c := make(chan requestResult)

	urls := []string{
		"https://www.reddit.com",
		"https://www.google.com",
		"https://nomadcoders.co",
	}

	var results = map[string]string{}

	for _, url := range urls {

		go hitURL(url, c)

	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
	/*
		results["hello"] = "hello"
		for url, result := range results {
			fmt.Println(url, result)
		}
	*/
}

func hitURL(url string, c chan<- requestResult) {
	fmt.Println("checking:", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url: url, status: "Failed"}
	} else {
		c <- requestResult{url: url, status: "Ok"}
	}
}
