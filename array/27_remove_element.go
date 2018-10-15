package main

import "fmt"

func main() {
	fmt.Println(removeElement1([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

//
// 遍历数组一如既往地联想到双指针法
//
func removeElement1(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast] // 快指针在前，慢指针在后，遇到不等的元素就放到慢指针的位置
			slow++
		}
	}
	return slow
}

// 类似 26 的解法
func removeElement2(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			if slow != fast {
				nums[slow] = nums[fast]
			}
			slow++
			fast++
			continue
		}
		fast++
	}
	return slow
}
