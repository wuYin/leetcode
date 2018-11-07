package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(7, 3)) // 7列 3行
}

// 感觉是个数学问题
func uniquePaths(m int, n int) int {
	paths := make([][]int, n)
	for i := range paths {
		paths[i] = make([]int, m)
	}

	for r := n - 1; r >= 0; r-- {
		paths[r][m-1] = 1
	}
	for c := m - 1; c >= 0; c-- {
		paths[n-1][c] = 1
	}

	for r := n - 1 - 1; r >= 0; r-- { // 从倒数第二行向前
		for c := m - 1 - 1; c >= 0; c-- { // 每行从后向前走
			paths[r][c] = paths[r][c+1] + paths[r+1][c]
		}
	}

	// fmt.Println(paths)
	// [28 21 15 10 6 3 1]
	// [7   6  5  4 3 2 1]
	// [1   1  1  1 1 1 1]
	return paths[0][0]
}
