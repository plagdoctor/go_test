package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	urls := []string{
		"https://www.reddit.com",
	}

	var results = make(map[string]string)
	// or var results = map[string]string{}

	results["hello"] = "hello"

	for _, url := range urls {
		result := "ok"
		err := hitURL(url)
		if err != nil {
			result = "failed"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("checking:", url)
	resp, err := http.Get(url)

	/*
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(data))
	*/

	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
