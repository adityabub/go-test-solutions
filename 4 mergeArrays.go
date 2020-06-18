//code
package main

import "fmt"

func mergeArrays(s1 []int, s2 []int, n1 int, n2 int, s3 []int) {
	var i, j, k = 0, 0, 0

	for i < n1 && j < n2 {
		if s1[i] < s2[j] {
			s3[k] = s1[i]
			k++
			i++
		} else {
			s3[k] = s2[j]
			k++
			j++
		}
	}

	for i < n1 {
		s3[k] = s1[i]
		k++
		i++
	}

	for j < n2 {
		s3[k] = s2[j]
		k++
		j++
	}

}

func main() {

	s1 := [4]int{1, 3, 4, 5}
	s2 := [4]int{2, 4, 6, 8}

	var s3 [8]int

	mergeArrays(s1[:], s2[:], 4, 4, s3[:])

	fmt.Println(s3)
}
