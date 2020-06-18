//code
package main

import "fmt"

func moveZeros(arr []int, n int) {
	var count int = 0
	for i := range arr {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}
	for count < n {
		arr[count] = 0
		count++
	}
}

func main() {
	s := [10]int{3, 5, 4, 0, 0, 2, 11, 8, 0, 6}
	moveZeros(s[:], 10)
	fmt.Println(s)
}
