//Iteration Code
package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var v = isPrime(19)
	if v {
		fmt.Println("It is a prime number")
	} else {
		fmt.Println("It is not a prime number")
	}
}
