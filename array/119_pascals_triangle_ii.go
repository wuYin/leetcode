package main

import "fmt"

func main() {
	fmt.Println(getRow(33))
}

func getRow(rowIndex int) []int {
	arrs := getTriangle(34)
	return arrs[rowIndex]
}

func getTriangle(numRows int) [][]int {
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
