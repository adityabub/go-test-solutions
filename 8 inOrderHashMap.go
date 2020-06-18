// What map generally does in Golang is sort the keys. Hence we cannot generate key-val pair in same order they were popuplated.
// What I have done is create 2 maps. First map m1 will map incremented int values(0, 1, 2, …) to given keys and second map m2 will map given keys to given values.
// When needed to be accessed in same order, we will range over m1 to get its values and use them as key to m2 to access in order in which it was populated.
// When only needed to be accessed by keys, we can use m2.
// This will take more space, generally ( <=2*normal_space) but will do the work.

package main

import "fmt"

var m1 map[int]int
var m2 map[int]int

func mapKeyValue(s1 []int, s2 []int) {
	m1 = make(map[int]int)
	m2 = make(map[int]int)
	count := 0
	for i := range s1 {
		k := s1[i]
		s := count
		v := s2[i]
		m1[s] = k
		m2[k] = v
		count++
	}

}

func sameOrder(i int) (int, int) {
	k := m1[i]
	v := m2[k]
	return k, v
}

func main() {

	s1 := [4]int{1, 4, 2, 3}
	s2 := [4]int{5, 8, 7, 6}

	mapKeyValue(s1[:], s2[:])

	fmt.Println(m1)
	fmt.Println(m2)

	for i := range m1 {
		k, v := sameOrder(i)
		fmt.Println(k, v)
	}
}
