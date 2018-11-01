package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
}

// 排序后累加偶数索引上的值，bingo
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i, num := range nums {
		if i%2 == 0 {
			sum += num
		}
	}
	return sum
}
