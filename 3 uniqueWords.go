//code
package main

import (
	"fmt"
	"regexp"
)

var m map[string]int

func printUniquedWords(str string) {
	var valid = regexp.MustCompile("[^A-Za-z0-9]")

	s := valid.Split(str, -1)

	m = make(map[string]int)

	totalCount := 0
	uniqueCount := 0

	for i := range s {
		if s[i] != "" {
			m[s[i]]++
			totalCount++
		}
	}
	for i := range m {
		if m[i] == 1 {
			uniqueCount++
		}
	}

	fmt.Println(totalCount, uniqueCount)
}

func main() {
	str := "ZBIO has come to campus for hiring today. ZBIO is a VC funded company in the enterprise infrastructure domain. This company is creating a platform to observe service to service communication"
	printUniquedWords(str)
}
