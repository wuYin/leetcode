package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combine(candidates, target)
}

func combine(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) <= 0 { // 递归到最后一个元素，递归结束
		return res
	}

	subSum := target - nums[0]

	switch {
	case subSum < 0: // 向后不存在更大的值，递归结束
		return res
	case subSum == 0: // 第一个数就是 target
		res = append(res, []int{nums[0]})
	case subSum > 0:
		remains := combine(nums, subSum) // 寻找所有的子和组合
		for _, v := range remains {
			way := append([]int{nums[0]}, v...)
			res = append(res, way)
		}
	}

	res = append(res, combine(nums[1:], target)...) // 向后查找其他组合，避免重复
	return res
}
