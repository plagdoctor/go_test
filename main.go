package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	c := make(chan string)
	people := []string{"nico", "jinyoung", "aaa", "bb", "talky"}

	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person + " is sexy"

}
