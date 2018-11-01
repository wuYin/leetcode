package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 0, 1})
	bestMoveZeroes([]int{1, 2, 3, 0, 0, 4})
}

// 辣鸡解法
func moveZeroes(nums []int) {
	n := len(nums)
	rear := n - 1
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			for j := i; j < rear; j++ { // 后部分全部往前挪动
				nums[j] = nums[j+1]
			}
			nums[rear] = 0
			rear--
			i-- // 原来0的位置上，不知道下一个值是不是0，所以 i--
		}
		if i >= rear {
			break
		}
	}
	fmt.Println(nums)
}

// 快慢指针
// fast 指针在前把所有不是 0 的元素挪动到 slow 指针上
func bestMoveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	fmt.Println(nums)
}
