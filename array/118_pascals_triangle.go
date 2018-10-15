package main

import "fmt"

func main() {
	fmt.Println(generate(0))
}

// 观察规律即可
func generate(numRows int) [][]int {
	var arrs [][]int
	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		arr[0], arr[i] = 1, 1
		for j := 1; j < i; j++ {
			arr[j] = arrs[i-1][j-1] + arrs[i-1][j]
		}
		arrs = append(arrs, arr)
	}
	return arrs
}
