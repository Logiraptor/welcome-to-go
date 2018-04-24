package main

import (
	"fmt"
)

func main() {
	for prime := range generatePrimes() {
		fmt.Println(prime)
	}
}

// GENERATE OMIT
func generatePrimes() chan int {
	result := make(chan int)

	go func() {
		for i := 2; ; i++ {
			if isPrime(i) {
				result <- i
			}
		}
	}()

	return result
}

// END GENERATE OMIT

// PRIME OMIT
func isPrime(x int) bool {
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// END PRIME OMIT
