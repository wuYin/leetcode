package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{
		// {1, 2, 3, 4},
		// {5, 6, 7, 8},
		// {9, 10, 11, 12},

		// {1, 2},
		// {3, 4},

		{2, 3, 4},
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
		{14, 15, 16},
	}))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return nil
	}
	if len(matrix) == 1 {
		return matrix[0]
	}
	var nums []int
	return order(matrix, 0, len(matrix[0])-1, 0, len(matrix)-1, nums)
}

func order(matrix [][]int, start, end int, up, down int, nums []int) []int {
	if start > end || up > down {
		return nums
	}
	// 向右走
	for i := start; i <= end; i++ {
		nums = append(nums, matrix[up][i])
	}

	// 向下走
	stop := true
	for i := up + 1; i <= down; i++ {
		nums = append(nums, matrix[i][end])
		stop = false
	}
	if stop {
		return nums // 无路可走
	}

	// 向左走
	stop = true
	for i := end - 1; i >= start; i-- {
		nums = append(nums, matrix[down][i])
		stop = false
	}
	if stop {
		return nums // 无路可走
	}

	// 向上走
	for i := down - 1; i > up; i-- {
		if end > 0 { // 单列的情况，下一趟不能往上走
			nums = append(nums, matrix[i][start])
		}
	}
	return order(matrix, start+1, end-1, up+1, down-1, nums)
}
