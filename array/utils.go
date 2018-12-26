package main

func min(vs ...int) int {
	if len(vs) == 0 {
		panic("func min args empty")
	}
	m := vs[0]
	for _, n := range vs[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func max(vs ...int) int {
	if len(vs) == 0 {
		panic("func min args empty")
	}
	m := vs[0]
	for _, n := range vs[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

// 递归版的二分查找虽高效，但调用传参时注意，nums 为空时 l,r 应为 0,-1，否则将 panic
func binarySearch(nums []int, l, r int, target int) int {
	if l > r {
		return -1
	}

	mid := (l + r) / 2
	switch {
	case target < nums[mid]:
		return binarySearch(nums, l, mid-1, target)
	case target > nums[mid]:
		return binarySearch(nums, mid+1, r, target)
	default:
		return mid
	}
}
