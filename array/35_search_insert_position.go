package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

// 二分查找
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		switch {
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] < target:
			left = mid + 1
		default:
			return mid
		}
	}
	return left
}
