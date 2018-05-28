package main

import (
	"time"
	"fmt"
	"sync"
)

func main() {
	var messages = make(chan string)
	go func() {
		defer close(messages)
		var wg = new(sync.WaitGroup)
		wg.Add(3)
		go sayAfter("Hello 1", time.Second, messages, wg)
		go sayAfter("Hello 2", 2*time.Second, messages, wg)
		go sayAfter("Hello 3", 3*time.Second, messages, wg)
		wg.Wait()
	}()

	for message := range messages {
		fmt.Println(message)
	}
}

func sayAfter(message string, d time.Duration, messages chan string, group *sync.WaitGroup) {
	time.Sleep(d)
	messages <- message
	group.Done()
}

// done OMIT
