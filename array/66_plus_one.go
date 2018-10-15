package main

import (
	"fmt"
)

func main() {
	// fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
}

// 若用 10 的次方做位数，聚合后再拆解，不确定长度的整数一定要考虑溢出，一旦溢出结果必定出错
func plusOne(digits []int) []int {
	n := len(digits)
	if digits[n-1]+1 < 10 {
		digits[n-1]++
	} else {
		for i := n - 1; i >= 0; i-- {
			if digits[i]+1 < 10 {
				digits[i]++
				break
			}
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		}
	}
	return digits
}
