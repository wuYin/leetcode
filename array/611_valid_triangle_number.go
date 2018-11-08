package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}

// 构成三角形的条件：a+b>c, a-b<c
// 两数之和的变种
func triangleNumber(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for _, num := range nums[j+1:] {
				if nums[j]-nums[i] < num && num < nums[i]+nums[j] {
					count++
				}
			}
		}
	}
	return count
}
