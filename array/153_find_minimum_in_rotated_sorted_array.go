package main

import "fmt"

func main() {
	// 在数组中查找原本 3 该在的位置，即是发生旋转的值（最小值）
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}

// 排序就能直接搞定，不过扩展的二分思路更有技巧
// 对二分查找顺序、查找完毕时的 l, r 变量的理解一定要深刻
func findMin(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return nums[0]
	}

	// 若在最后一个元素上旋转，数组依旧有序
	if nums[n-1] > nums[0] {
		return nums[0]
	}

	target := nums[0]
	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		// 不断在无序区找最小值
		switch {
		case nums[mid] >= target: // 中间值比目标大，前半部分有序，继续在后部分无序区找最小值
			l = mid + 1
		case nums[mid] < target: // 中间值比目标小，前半部分无序
			r = mid
		}
	}
	return nums[r]
}
