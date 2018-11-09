package main

import "fmt"

func main() {
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0})) // 1
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4})) // 4
}

// 规律性很强
// 对长度为 n 的排列 [0,..., n-1] 要敏感一些
func maxChunksToSorted(arr []int) int {
	max := arr[0]
	count := 0
	for i, n := range arr {
		if n > max {
			max = n
		}
		if max == i {
			count++ // 子区间最大值是当前索引，则为独立空间
		}
	}
	return count
}
