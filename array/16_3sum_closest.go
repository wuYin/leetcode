package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest1([]int{-1, 2, 1, -4}, 1))                 // 2
	fmt.Println(threeSumClosest2([]int{0, 2, 1, -3}, 1))                  // 0
	fmt.Println(threeSumClosest2([]int{1, 2, 4, 8, 16, 32, 64, 128}, 84)) // 84
}

// 同样也是指针法
func threeSumClosest2(nums []int, target int) int {
	sort.Ints(nums)
	minSum := nums[0] + nums[1] + nums[2]
	minAbs := abs(minSum - target)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		head, tail := i+1, len(nums)-1
		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			newAbs := abs(sum - target)
			if newAbs < minAbs { // 找到了更接近的和
				minAbs = newAbs
				minSum = sum
			}
			if sum < target { // 值小了，往后移，可能更接近
				head++
			} else {
				tail-- // 值大了，往前移动
			}
		}
	}
	return minSum
}

// 暴力三层遍历，求解最接近的和
func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	minAbs := abs(nums[0] + nums[1] + nums[2] - target)
	minSum := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) <= minAbs {
					minSum = sum
					minAbs = abs(sum - target)
				}
			}
		}
	}
	return minSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
