package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)) // 5
}

func findKthLargest(nums []int, k int) int {
	maxHeap := NewMaxHeap(nums)
	for k > 1 {
		maxHeap.Pop()
		k--
	}
	return maxHeap.Pop()
}
