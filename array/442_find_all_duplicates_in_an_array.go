package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})) // [2, 3]
}

// 多么适合用哈希表的题...
// 可惜不能使用额外空间，复杂度 O(N)，可能要求在数组原地操作、一次遍历解决
// 解法和下方的 448 很像，充分利用 1~n 的已知条件，十分巧妙
func findDuplicates(nums []int) []int {
	var dups []int
	for _, num := range nums {
		if num < 0 {
			num = -num
		}

		n := num - 1
		if nums[n] < 0 { // 指向的位置以为负值，则重复
			dups = append(dups, num)
			continue
		}
		nums[n] = -nums[n]
	}
	return dups
}
