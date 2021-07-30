package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "jinyoung"}

	for _, person := range people {
		go isSexy(person, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true

}
