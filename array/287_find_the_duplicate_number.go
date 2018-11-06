package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2})) // 2
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2})) // 3
}

// 和 442、448 就比较像，可惜题目要求数组
// 不可修改：不能排序、不能取负值
// 空间复杂度O(1)：不能使用 O(N) 的哈希表
// 时间复杂度O(N^2)
func findDuplicate(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	slow := nums[0]
	fast := nums[slow]
	for slow != fast { // 题给必有一个重复值
		slow = nums[slow]
		fast = nums[nums[fast]] // slow走一步，fast走两步，fast多绕一圈追上了slow，二者指向 loop 的入口
	}
	fmt.Println(slow, fast)

	fast = 0 // 从头出发去找 loop 的入口
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return slow
}
