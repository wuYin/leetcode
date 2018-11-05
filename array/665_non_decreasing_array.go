package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3})) // false
	fmt.Println(checkPossibility([]int{1, 4, 2, 3})) // true
}

func checkPossibility(nums []int) bool {
	smallCount := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if smallCount == 1 {
				return false
			}

			// 巧妙一步：如果 i-2 的数比 i 的大，就向后延续大数的影响范围
			// 3 4 2 3	// 3 4 4 3
			// 1 4 2 3	// 1 4 2 3
			if i >= 2 && nums[i-2] >= nums[i] {
				nums[i] = nums[i-1]
			}
			// fmt.Println(nums)
			smallCount++
		}
	}
	return true
}
