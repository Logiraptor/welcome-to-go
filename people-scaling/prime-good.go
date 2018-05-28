package main

// START OMIT
func isPrime(x int) bool {
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
// END OMIT
