package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{1, 0, 100}))          // 0
	fmt.Println(maximumProduct([]int{-4, -3, -2, -1, 60})) // 720
}

// 注意各种逻辑判断即可
func maximumProduct(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	sort.Ints(nums)
	maxNegMulti := nums[0] * nums[1] // 最小的两个数之积

	multi := 1
	for i := len(nums) - 1; i >= 0 && i > len(nums)-1-3; i-- {
		multi *= nums[i]
	}

	if maxNegMulti > 0 && multi < maxNegMulti*nums[len(nums)-1] {
		return maxNegMulti * nums[len(nums)-1] // 最大负数积乘最大正数
	}

	return multi
}
