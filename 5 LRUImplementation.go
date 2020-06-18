//code
package main

import (
	"container/list"
	"fmt"
)

// FixedSize
var size = 4

// Add recent page
func add(l *list.List, x int) {

	// Check if x is present in cache
	e := l.Front()
	f := 0
	for ; e != nil; e = e.Next() {
		if e.Value == x {
			f = 1
			break
		}
	}

	// not present in cache
	if f == 0 {
		if l.Len() == size {
			// Cache is full. Remove last and add to front
			l.Remove(l.Back())
			l.PushFront(x)
		} else {
			// Add to front.
			l.PushFront(x)
		}
	} else {
		// x is present. Move to front.
		l.MoveToFront(e)
	}
}

// Show all the pages currently in cache
func show(l *list.List) {
	// print all elements
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// return current length of LRU
func len(l *list.List) int {
	return l.Len()
}

func main() {

	// Queue as doubly linked list. AddNewPage: Front, RemoveOldPage: Back.
	l := list.New()

	add(l, 1)
	add(l, 2)
	add(l, 3)

	//fmt.Println(len(l))

	add(l, 1)
	add(l, 4)
	add(l, 5)

	show(l)
}
