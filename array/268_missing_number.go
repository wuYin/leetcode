package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(bestMissingNumber([]int{3, 0, 1}))
	fmt.Println(bestMissingNumber([]int{0}))
}

// 让我想起了小学时候学的急转弯...
// 如何计算一个人的体积：把人放入一桶水中，计算桶前后的体积差即可...
// 如果题目给出的已知条件十分地有规律，那解决方法肯定是有规律的，尝试着找出来
func bestMissingNumber(nums []int) int {
	curSum := 0
	for i := range nums {
		curSum += nums[i]
	}
	n := len(nums)
	fullSum := n * (n + 1) / 2
	return fullSum - curSum
}

// 辣鸡解法
func missingNumber(nums []int) int {
	sort.Ints(nums) // lgN
	for i := 0; i < len(nums); i++ { // N
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
