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
