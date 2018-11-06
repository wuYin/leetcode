package main

import (
	"fmt"
)

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

// 考虑使用递归实现，到达 grid[i][j] 的最小路径为到达 grid[i-1][j] 和 grid[i][j-1] 的最小值
// 但递归不保存任何已计算出的结果，效率不高，类比斐波那契数列的递归实现
// 使用简单的动态规划，即保存中间计算结果
func minPathSum(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}

	row, column := len(grid), len(grid[0])
	steps := make([][]int, row) // 存储到达 [i,j] 的最少步数
	for i := range steps {
		steps[i] = make([]int, column)
	}
	steps[0][0] = grid[0][0]

	// 只能向下或向左移动，先求解简单值
	// 首行
	for c := 1; c < column; c++ {
		steps[0][c] = grid[0][c] + steps[0][c-1]
	}
	// 首列
	for r := 1; r < row; r++ {
		steps[r][0] = grid[r][0] + steps[r-1][0]
	}

	// 再根据简单值，逐步求解复杂值
	// 从 [1][1] 开始走到右下顶点
	for r := 1; r < row; r++ {
		for c := 1; c < column; c++ {
			// 从上方，左方选最小的权值
			steps[r][c] = min(steps[r-1][c], steps[r][c-1]) + grid[r][c]
		}
	}

	return steps[row-1][column-1] // 累加到右下角的值是最小权值
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
