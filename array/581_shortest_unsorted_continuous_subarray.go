package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{2, 1}))
}

// 先排序，再双指针 ok
func findUnsortedSubarray(nums []int) int {
	backup := make([]int, len(nums))
	copy(backup, nums)

	sort.Ints(nums)
	start, end := -1, -1
	for i, num := range backup {
		if nums[i] != num {
			start = i
			break
		}
	}
	if start == -1 {
		return 0
	}

	for i := len(backup) - 1; i >= 0; i-- {
		if nums[i] != backup[i] {
			end = i
			break
		}
	}

	return end - start + 1
}
