package main

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap([]int{1, 4, 6, 7, 3})
	fmt.Println(*h) // [1 3 6 7 4] // ok
	h.Push(2)
	fmt.Println(*h)      // [1 3 2 7 4 6] // ok
	fmt.Println(h.Pop()) // 1
	fmt.Println(*h)      // [2 3 6 7 4] // ok
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap([]int{1, 2, 6, 5, 9})
	fmt.Println(*h) // [9 5 6 1 2]
	h.Push(8)
	fmt.Println(*h) // [9 5 8 1 2 6]
	fmt.Println(h.Pop())
	fmt.Println(*h) // [8 5 6 1 2]
}
