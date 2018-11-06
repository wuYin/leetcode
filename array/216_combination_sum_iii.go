package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(4, 24)) // [[8 9 6 1] [8 9 5 2] [7 9 6 2] [8 9 4 3] [7 9 5 3] [7 8 6 3] [6 9 5 4] [7 8 5 4]]
	fmt.Println(combinationSum3(3, 7))  // [[2 4 1]]
	fmt.Println(combinationSum3(3, 9))  // [[2 6 1] [3 5 1] [3 4 2]]
}

// 调了一小时才调出来的递归写法阿西吧
// 递归函数必定有退出条件，先分析递归执行到最后如何退出，再去写一般的递归传参逻辑
// 本题是 k 递减到 2 时求两数之和再退出
func combinationSum3(k int, n int) [][]int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var res [][]int

	switch {
	case k == 0:
		return nil
	case k == 1:
		if 1 <= n && n <= 9 {
			return [][]int{{n}}
		}
	default:
		res = append(res, combineSum(k, n, nums)...)
	}

	return res
}

func combineSum(k, target int, nums []int) [][]int {
	// fmt.Printf("%v, %v, %v\n", k, target, nums) // 输出入参来调试递归
	if k <= 0 || len(nums) <= 0 || target <= 0 {
		return nil
	}

	// 递归退出条件：k 自减到 2
	if k == 2 {
		return combineTwoSum(nums, target)
	}

	var res [][]int
	for i := 0; i < len(nums); i++ {
		// 关键的递归步骤：对于当前值 nums[i]，再从 nums[i+1:] 余下数组中找和为 target-nums[i] 的 k-1 个数
		remains := combineSum(k-1, target-nums[i], nums[i+1:])
		if len(remains) <= 0 { // 没找到
			continue
		}
		for _, remain := range remains { // 找到了，就把 nums[i] 匹配进去
			remain = append(remain, nums[i])
			res = append(res, remain)
		}
	}

	return res
}

// nums 中所有和为 target 的两数之和
func combineTwoSum(nums []int, target int) [][]int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}

	var res [][]int
	for _, n := range nums {
		if n >= target {
			break
		}
		diff := target - n
		if _, ok := m[diff]; ok && n != diff {
			res = append(res, []int{n, diff})
		}
	}
	return res[:len(res)/2] // 去重
}
