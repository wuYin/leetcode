package main

import "fmt"

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},

		// {-1},
		// {2, 3},
		// {1, -1, -1},
	}
	fmt.Println(minimumTotal(triangle)) // 2 3 5 1
}

// 和 64 题最小路径和很像，都是用简单的动态规划实现
func minimumTotal(triangle [][]int) int {
	r := len(triangle)
	if r <= 0 {
		return 0
	}

	// 向下构建
	steps := make([][]int, r)
	for i := range triangle {
		steps[i] = make([]int, i+1)
	}
	copy(steps, triangle)

	for i := 1; i < r; i++ {
		for j := 0; j <= i; j++ { // 第 N 行有 N 个元素
			steps[i][j] += min(steps[i-1], j)
		}
	}

	minStep := steps[r-1][0]
	for _, step := range steps[r-1] {
		if minStep > step {
			minStep = step
		}
	}
	return minStep
}

// 获取某行指定位置及其左右的最小值
func min(nums []int, c int) (minNum int) {
	n := len(nums)
	l, mid := c-1, c

	inited := false
	if 0 <= l && l <= n-1 {
		inited = true
		minNum = nums[l]
	}
	if 0 <= mid && mid <= n-1 {
		if inited {
			if minNum > nums[mid] {
				minNum = nums[mid]
			}
		} else {
			inited = true
			minNum = nums[mid]
		}
	}
	return
}
