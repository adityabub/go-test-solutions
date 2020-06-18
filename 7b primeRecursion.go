// Recursion Code
package main

import "fmt"

func isPrime(n int, k int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%k == 0 {
		return false
	}
	if k*k > n {
		return true
	}
	return isPrime(n, k+1)
}

func main() {
	var v = isPrime(23, 2)
	if v {
		fmt.Println("It is a prime number")
	} else {
		fmt.Println("It is not a prime number")
	}
}
