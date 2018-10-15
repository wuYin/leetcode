package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))  // [[-1 -1 2] [-1 0 1]]
	fmt.Println(badTwoSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))       // [[-2 0 2]]
}

//
// 双指针法
// 遍历元素，取剩余元素的 head 和 tail 指针，二者向中间移动，如果和为该元素的相反数，则找到三元组
// 注意剔除连续相等的数，不然会导致三元组重复，如 [-2,0,0,2,2]
//
func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums); i++ {
		remain := 0 - nums[i]
		head, tail := i+1, len(nums)-1

		if nums[i] > 0 { // 优化：之后全是正数，不可能组成三元组
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 剔除遍历到连续相等的数
			continue
		}

		for head < tail {
			sum := nums[head] + nums[tail]
			switch {
			case sum == remain: // 找到三元组
				res = append(res, []int{nums[i], nums[head], nums[tail]})
				for head < tail && nums[head] == nums[head+1] { // 剔除待选数字中连续相等的数
					head++
				}
				for head < tail && nums[tail] == nums[tail-1] {
					tail--
				}
				head++
				tail--
			case sum < remain: // 往后走
				head++
			case sum > remain: // 往前走
				tail--
			}
		}
	}
	return res
}

//
// twoSum 的思路，不好
//
func badTwoSum(nums []int) [][]int {
	// 避开全是 0 的 case	 // ugly
	if len(nums) >= 3 {
		allZero := true
		for _, num := range nums {
			if num != 0 {
				allZero = false
			}
		}
		if allZero {
			return [][]int{{0, 0, 0}}
		}
	}

	n := len(nums)
	num2index := make(map[int]int, n)
	for i, num := range nums {
		num2index[num] = i
	}

	// 获取三元组
	var res [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			remain := 0 - (nums[i] + nums[j])
			if k, ok := num2index[remain]; ok && j != k && i != k {
				res = append(res, []int{nums[i], nums[j], remain})
			}
		}
	}

	// 剔除重复的三元组
	m := make(map[string][]int)
	for i := range res {
		sort.Ints(res[i])
		m[intStr(res[i])] = res[i]
	}

	var arrs [][]int
	for _, arr := range m {
		arrs = append(arrs, arr)
	}
	return arrs
}

// trick	// 使得整数数组能做 map 的 key
func intStr(nums []int) string {
	str := ""
	for _, num := range nums {
		str += fmt.Sprintf("%d", num)
	}
	return str
}
