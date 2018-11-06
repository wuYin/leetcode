package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4})) // [24 12 8 6]
	fmt.Println(productExceptSelf([]int{2, 4, 5, 7})) // [140 70 56 40]
}

// 其余部分积 = 左子数组积 * 右子数组积
// 遍历方式很巧妙
// 可惜 Golang 实现执行时间貌似都过不了 50000 多个[1, -1, ...]的那个 case
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1] // 把 nums[i] 左测子数组的积存在 res[i] 中
	}
	// fmt.Println(res) // [1 2 8 40]

	back := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= back // res[i] 再乘右侧积
		back *= nums[i]
	}
	return res
}
