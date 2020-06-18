//code
package main

import "fmt"

var m map[int]int

func mapKeyValue(s1 []int, s2 []int) {
	m = make(map[int]int)
	for i := range s1 {
		k := s1[i]
		v := s2[i]
		m[k] = v
	}

}

func main() {

	s1 := [4]int{1, 3, 4, 5}
	s2 := [4]int{2, 4, 6, 8}

	mapKeyValue(s1[:], s2[:])

	fmt.Println(m)
}
