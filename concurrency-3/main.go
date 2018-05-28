package main

import (
	"time"
	"fmt"
)

func main() {
	var messages = make(chan string)
	go sayAfter("Hello 1", time.Second, messages)
	go sayAfter("Hello 2", 2*time.Second, messages)
	go sayAfter("Hello 3", 3*time.Second, messages)

	for message := range messages {
		fmt.Println(message)
	}
}

func sayAfter(message string, d time.Duration, messages chan string) {
	time.Sleep(d)
	messages <- message
}
