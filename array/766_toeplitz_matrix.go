package main

import "fmt"

func main() {
	fmt.Println(isToeplitzMatrix([][]int{
		// {1, 2, 3, 4},
		// {5, 1, 2, 3},
		// {9, 5, 1, 2},

		// {1, 2},
		// {2, 2},
		{36, 59, 71, 15, 26, 82, 87},
		{56, 36, 59, 71, 15, 26, 82},
		{15, 0, 36, 59, 71, 15, 26},
	}))
}

// 遍历第一行，对角线检测
// 遍历第一列，对角线检测
// 注意 2 个边界可以少检测 2 次
func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) <= 1 || len(matrix[0]) <= 1 {
		return true
	}

	rows, column := len(matrix), len(matrix[0])
	// 横向检测
	for i := 0; i < column-1; i++ {
		x, y := 0, i
		for x+1 < rows && y+1 < column {
			if matrix[x][y] != matrix[x+1][y+1] {
				return false
			}
			x++
			y++
		}
	}
	// 竖向检测
	for i := 1; i < rows-1; i++ {
		x, y := i, 0
		for x+1 < rows && y+1 < column {
			if matrix[x][y] != matrix[x+1][y+1] {
				return false
			}
			x++
			y++
		}
	}
	return true
}
