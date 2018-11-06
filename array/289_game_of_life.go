package main

import "fmt"

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board) // [[0 0 0] [1 0 1] [0 1 1] [0 1 0]]
}

// 题目描述长长长...
func gameOfLife(board [][]int) {
	backup := make([][]int, len(board)) // 生死同时发生
	for i := range board {
		backup[i] = make([]int, len(board[0]))
		copy(backup[i], board[i])
	}

	for i := range backup {
		for j := range backup[i] {
			counts := countAround(backup, i, j)
			if backup[i][j] == 1 && (counts < 2 || counts > 3) { // 活细胞狗带
				board[i][j] = 0
			}

			if backup[i][j] == 0 && counts == 3 { // 死细胞复活
				board[i][j] = 1
			}
		}
	}
	fmt.Println(board)
}

func countAround(board [][]int, i, j int) int {

	r, c := len(board)-1, len(board[0])-1
	count := 0
	if i > 0 {
		if board[i-1][j] == 1 { // 上
			count++
		}
		if j > 0 && board[i-1][j-1] == 1 { // 左上
			count++
		}
	}
	if i < r {
		if board[i+1][j] == 1 { // 下
			count++
		}
		if j > 0 && board[i+1][j-1] == 1 { // 左下
			count++
		}
	}

	if j > 0 {
		if board[i][j-1] == 1 { // 左
			count++
		}
	}
	if j < c {
		if board[i][j+1] == 1 { // 右
			count++
		}
		if i > 0 && board[i-1][j+1] == 1 { // 右上
			count++
		}
		if i < r && board[i+1][j+1] == 1 { // 右下
			count++
		}
	}
	return count
}
