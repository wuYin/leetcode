package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{3, 4, 3, 2, 1}))       // 1
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})) // 5 / 1
}

// 二分查找的变种
func findPeakElement(nums []int) int {
	return findPeak(nums, 0, len(nums))
}

func findPeak(nums []int, l, r int) int {
	mid := (l + r) / 2

	inLeft, inRight := false, false
	if mid > 0 && nums[mid-1] > nums[mid] { // 峰值在左侧
		inLeft = true
	}
	if mid < len(nums)-1 && nums[mid] < nums[mid+1] { // 在右侧
		inRight = true
	}

	switch {
	case inLeft:
		return findPeak(nums, l, mid-1)
	case inRight:
		return findPeak(nums, mid+1, r)
	default:
		return mid
	}
}
