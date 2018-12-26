package main

import "fmt"

func main() {
	fmt.Println(searchRange1([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
	fmt.Println(searchRange2([]int{1}, 1))                 // [0, 0]
}

// 直接二分查找
// O(N) // not ok
func searchRange1(nums []int, target int) []int {
	n := len(nums)
	i := binarySearch(nums, 0, n-1, target)
	if i == -1 {
		return []int{-1, -1}
	}

	// 找到后向两侧延伸
	l, r := i, i
	for j := i - 1; j >= 0; j-- {
		if nums[j] == target {
			l = j
		}
	}
	for j := i + 1; j < n; j++ {
		if nums[j] == target {
			r = j
		}
	}
	return []int{l, r}
}

// 改进的二分搜索
func searchRange2(nums []int, target int) []int {
	l := edgeBinSearch(nums, true, target)
	r := edgeBinSearch(nums, false, target)
	return []int{l, r}
}

// 普通二分搜索在匹配到 target 时直接 return
// 在本题搜索时在匹配到 target 之后依旧向边缘走当做没匹配到，注意 2 个边界条件
// O(logN) // ok
func edgeBinSearch(nums []int, leftest bool, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case target < nums[mid]:
			r = mid - 1
		case target > nums[mid]:
			l = mid + 1
		default:
			if leftest {
				if mid == 0 || nums[mid] > nums[mid-1] { // 不再继续向左走的 2 个边界条件
					return mid
				}
				r = mid - 1 // 继续在左侧找
			} else {
				if mid == n-1 || nums[mid] < nums[mid+1] {
					return mid
				}
				l = mid + 1 // 继续在右侧找
			}
		}
	}
	return -1
}
