package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

// 简单的逻辑...
func sortArrayByParity(A []int) []int {
	isOdd := func(x int) bool {
		return x%2 == 1
	}
	odds := make([]int, 0, len(A))
	evens := make([]int, 0, len(A))
	for _, a := range A {
		if isOdd(a) {
			odds = append(odds, a)
			continue
		}
		evens = append(evens, a)
	}
	return append(evens, odds...)
}
