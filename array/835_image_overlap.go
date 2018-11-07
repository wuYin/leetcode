package main

import "fmt"

func main() {
	a := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	b := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}
	fmt.Println(largestOverlap(a, b)) // 3
}

// n <= 30，暴力遍历可接受
func largestOverlap(A [][]int, B [][]int) int {
	max := 0
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j ++ {

			countA, countB := 0, 0
			for r := i; r < n; r++ {
				for c := j; c < n; c++ {

					preR, preC := r-i, c-j
					if A[r][c] == 1 && B[preR][preC] == 1 { // 重复且为 1 的位置计数
						countA += 1
					}
					if B[r][c] == 1 && A[preR][preC] == 1 {
						countB += 1
					}
				}
			}

			if countA > max {
				max = countA
			}
			if countB > max {
				max = countB
			}
		}
	}
	return max
}
