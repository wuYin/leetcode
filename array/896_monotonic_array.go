package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3})) // true
	fmt.Println(isMonotonic([]int{1}))          // true
}

func isMonotonic(A []int) bool {
	isUp, isDown := true, true
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			isUp = false
		}
		if A[i] < A[i+1] {
			isDown = false
		}
		if !isUp && !isDown { // 优化：提前结束遍历
			return false
		}
	}
	return isUp || isDown
}
