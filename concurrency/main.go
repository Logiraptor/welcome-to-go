package main

import (
	"time"
	"fmt"
)

func main() {
	sayAfter("Hello 1", time.Second)
	sayAfter("Hello 2", 2*time.Second)
	sayAfter("Hello 3", 3*time.Second)
}

func sayAfter(message string, d time.Duration) {
	time.Sleep(d)
	fmt.Println(message)
}
