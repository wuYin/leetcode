package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
}

// 类二分搜索
// 最左边数 < 中间数则左侧有序，最右边数 > 中间数则右侧有序
// 在缩小搜索区域时，一直只在确定的有序区域内查找
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case nums[mid] == target: // bingo
			return mid
		case nums[l] <= nums[mid]: // 左侧有序
			if nums[l] <= target && target < nums[mid] { // 保证 target 一定在有序的左侧内
				r = mid - 1
			} else {
				l = mid + 1
			}
		case nums[mid] <= nums[r]: // 右侧有序
			if nums[mid] < target && target <= nums[r] { // 保证 target 一定在有序的右侧内
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
