package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

//
// 针对有序数组，双指针法是十分常见且有用的
//
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums)-1 {
		if nums[fast] != nums[fast+1] { // 相邻的数不相等
			slow++
			fast++
			nums[slow] = nums[fast] // 将最新的新数存储到慢指针的位置
			continue
		}
		fast++
	}
	return slow + 1
}
