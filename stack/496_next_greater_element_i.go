package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))     // [-1 3 -1]
	fmt.Println(bestNextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})) // [-1 3 -1]
}

// 使用递减顺序栈，很 ok
func bestNextGreaterElement(findNums []int, nums []int) []int {
	m := make(map[int]int)
	var s Stack // 递减序列栈
	for _, num := range nums {
		for !s.IsEmpty() && s.Peek() < num { // 将 nums 中的数依次进栈，一直向下找比栈顶值还大的数
			m[s.Pop()] = num // 向上最接近的数
		}
		s.Push(num)
	}

	for i, num := range findNums {
		if greater, ok := m[num]; ok {
			findNums[i] = greater
		} else {
			findNums[i] = -1
		}
	}
	return findNums
}

// 干脆直接的遍历
func nextGreaterElement(findNums []int, nums []int) []int {
	n := len(nums)
	res := make([]int, 0, len(findNums))
	for _, target := range findNums {
		for i, num := range nums {
			if num != target {
				continue
			}

			// 直接向后找第一个比 target 大的数
			found := false
			for j := i + 1; j < n; j++ {
				if nums[j] > target {
					found = true
					res = append(res, nums[j])
					break
				}
			}
			if !found {
				res = append(res, -1)
			}
		}
	}
	return res
}
