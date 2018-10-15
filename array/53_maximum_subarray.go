package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-3, 4, -1, 3, -5}))
	fmt.Println(maxSubArray([]int{-2, -1}))
}

//
// 跨步遍历的感觉
//
func maxSubArray(nums []int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	sum, maxSum := 0, 0
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] > 0:
			sum += nums[i]
		case nums[i] < 0:
			if sum+nums[i] > 0 {
				sum += nums[i]
			} else {
				sum = 0
			}
		}

		if maxSum < sum {
			maxSum = sum
		}
	}
	if max < 0 && maxSum > max {
		return max
	}
	return maxSum
}
