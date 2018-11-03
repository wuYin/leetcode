package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0})) // 1
	fmt.Println(dominantIndex([]int{1, 2, 3, 4})) // 1
}

// 辣鸡解法
func dominantIndex(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	sort.Ints(nums)
	if nums[n-1] >= 2*nums[n-2] {
		return m[nums[n-1]]
	}
	return -1
}
