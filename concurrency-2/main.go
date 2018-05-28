package main

import (
	"time"
	"fmt"
)

func main() {
	go sayAfter("Hello 1", time.Second)
	go sayAfter("Hello 2", 2*time.Second)
	go sayAfter("Hello 3", 3*time.Second)
	time.Sleep(time.Second * 4)
}

func sayAfter(message string, d time.Duration) {
	time.Sleep(d)
	fmt.Println(message)
}
