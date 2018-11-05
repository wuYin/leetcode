package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))      // 3
	fmt.Println(pivotIndex([]int{1, 2, 3}))               // -1
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, 0})) // 2
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, 0, -1})) // 2
}

// 左右两端和相等，双指针向中间遍历
// 这里的指针不是索引，而是递增和、递减和
func pivotIndex(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return -1
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	leftSum := 0
	for i := 0; i < n; i++ {
		if leftSum == sum-nums[i]-leftSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
