package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	for prime := range generatePrimes(ctx) {
		fmt.Println(prime)
	}
}

// GENERATE OMIT
func generatePrimes(ctx context.Context) chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		for i := 2; ; i++ {
			if isPrime(i) {
				result <- i
			}

			select {
			case <-ctx.Done():
				return
			default:
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
