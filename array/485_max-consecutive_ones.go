package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{}))                 // 0
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})) // 3
}

func findMaxConsecutiveOnes(nums []int) int {
	n := len(nums)
	count, maxCount := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			count = 1
			for j := i + 1; j < n; j++ { // 向后找连续的 1
				if nums[j] == 0 {
					break
				}
				count++
				i++
			}
			maxCount = max(count, maxCount)
		}
	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
