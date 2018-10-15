package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}

//
// 与有序数组相关的搜索，应当优先联想到二分查找的思路
// 旋转的本质是改变部分序列的排序规则，应当发现以下规律：
// 最左边数 < 中间数则左侧有序，最右边数 > 中间数则右侧有序，再利用二分查找的思路即可
//
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < nums[right]: // 右侧有序
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1 // 在右侧有序区
			} else {
				right = mid - 1 // 进入左侧无序区查找
			}
		case nums[mid] >= nums[left]: // 左侧有序
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1 // 在左侧有序区
			} else {
				left = mid + 1 // 进入右侧无序区查找
			}
		}
	}
	return -1
}
