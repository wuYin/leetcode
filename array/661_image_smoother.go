package main

import "fmt"

func main() {
	m := [][]int{
		// {1, 1, 1},
		// {1, 0, 1},
		// {1, 1, 1},

		{2, 3, 4},
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
		{14, 15, 16},
		// [[4 4 5] [5 6 6] [8 9 9] [11 12 12] [13 13 14]]
	}
	fmt.Println(imageSmoother(m))
}

// 免不了一堆判断，取周边和再求平均值
func imageSmoother(M [][]int) [][]int {
	if len(M) <= 0 || len(M[0]) <= 0 {
		return nil
	}
	r, c := len(M), len(M[0])
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			sum := M[i][j]
			k := 1
			if i > 0 {
				sum += M[i-1][j]
				k++
			}
			if i < r-1 {
				sum += M[i+1][j]
				k++
			}
			if j > 0 {
				sum += M[i][j-1]
				k++
			}
			if j < c-1 {
				sum += M[i][j+1]
				k++
			}
			if i > 0 && j > 0 {
				sum += M[i-1][j-1]
				k++
			}
			if i > 0 && j < c-1 {
				sum += M[i-1][j+1]
				k++
			}
			if i < r-1 && j > 0 {
				sum += M[i+1][j-1]
				k++
			}
			if i < r-1 && j < c-1 {
				sum += M[i+1][j+1]
				k++
			}
			res[i][j] = sum / k
		}
	}
	return res
}
