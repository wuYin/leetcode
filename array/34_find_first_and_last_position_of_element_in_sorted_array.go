package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

//
// 二分查找
// 注意找到一个数后向前、向后遍历找到边界值即可
//
func searchRange(nums []int, target int) []int {
	loc := binarySearch(nums, target)
	if loc == -1 {
		return []int{-1, - 1}
	}
	left, right := loc, loc
	for left > 0 {
		if nums[left] != nums[left-1] {
			break
		}
		left--
	}
	for right < len(nums)-1 {
		if nums[right] != nums[right+1] {
			break
		}
		right++
	}
	return []int{left, right}
}

func binarySearch(nums []int, target int) int {
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
	return -1
}
