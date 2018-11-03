package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

// 做法比较巧妙
// 充分利用 [1, .., n] 的已知条件
// 为了标记 N 出现过，可以把 nums[N-1] 置为负数
// 最后遍历，没有被标记过的位置 index 即是未出现的数
// 巧妙地利用了标记位置和索引的关系
func findDisappearedNumbers(nums []int) []int {
	for _, n := range nums {
		if n < 0 {
			n = -n // 为了 O(1) 的空间复杂度，避免数组越界
		}
		if nums[n-1] > 0 {
			nums[n-1] = -nums[n-1]
		}
	}

	var vanishNums []int
	for i, n := range nums {
		if n > 0 {
			vanishNums = append(vanishNums, i+1)
		}
	}
	return vanishNums
}
