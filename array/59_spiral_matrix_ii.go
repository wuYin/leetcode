package main

import "fmt"

func main() {
	matrix := generateMatrix(3)
	for _, row := range matrix {
		for _, num := range row {
			fmt.Printf("%d\t", num) // it's ok
		}
		fmt.Println()
	}
}

// 感觉没啥规律可言啊，只能自己绕圈圈走呗
// 注意换方向时当前行索引的自增、自减
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	r, c := 0, 0 // 当前行、当前列
	for num := 1; num <= n*n; {
		// 从左向右
		for c < n && matrix[r][c] == 0 {
			matrix[r][c] = num
			num++
			c++
		}
		c--
		r++

		// 从上到下
		for r < n && matrix[r][c] == 0 {
			matrix[r][c] = num
			num++
			r++
		}
		r--
		c--

		// 从右到左
		for c >= 0 && matrix[r][c] == 0 {
			matrix[r][c] = num
			num++
			c--
		}
		r--
		c++

		// 从下到上
		for r >= 0 && matrix[r][c] == 0 {
			matrix[r][c] = num
			num++
			r--
		}
		r++
		c++
	}
	return matrix
}
