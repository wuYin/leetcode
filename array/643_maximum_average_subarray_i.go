package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{0, 1, 1, 3, 3}, 4))
}

// 固定长度的子序列求和，像尺子上的蚯蚓一样向前挪...
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	for i := k; i < len(nums); i++ {
		// 挪动...
		sum += nums[i]
		sum -= nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}
