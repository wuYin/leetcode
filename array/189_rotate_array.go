package main

import "fmt"

func main() {
	nums := []int{-1, -100, 3, 99}
	rotate2(nums, 2)
	fmt.Println(nums)
}

// 每次都把最后一个数挪到第一位来即可
func rotate1(nums []int, k int) {
	n := len(nums)
	for k > 0 {
		last := nums[n-1]
		for i := n - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = last
		k--
	}
}

func rotate2(nums []int, k int) {
	n := len(nums)
	step := n - k

}
