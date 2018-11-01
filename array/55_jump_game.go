package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}, true)) // true
	fmt.Println(jump([]int{3, 2, 1, 0, 4}, true)) // false
	fmt.Println(jump([]int{1, 2, 3}, true))       // true
	fmt.Println(jump([]int{0, 1}, true))          // false
	fmt.Println(jump([]int{0, 2, 3}, true))       // false
	fmt.Println(jump([]int{2, 0, 0}, true))       // true
	fmt.Println(jump([]int{2, 0, 1, 0, 1}, true)) // false
}

func canJump(nums []int) bool {
	return jump(nums, true)
}

// 要想达到最后一位，数组的前 n-1 个成员必须 >= 直达数组的各成员
// 如 [2 3 1 1] >=  [4 3 2 1]
//    [2, 3, 1] >= [3 2 1]
//    [2, 3] >= [2 1]
//    [2] >= [1]
// 效率比较低的递归
func jump(nums []int, ok bool) bool {
	n := len(nums)
	if !ok {
		return false
	}

	// 处理遇到 0 的极端情况
	for i := n - 1 - 1; i >= 0; i-- {
		if nums[i] == 0 {
			direct := 1
			for j := i - 1; j >= 0; j-- {
				// fmt.Println(nums[j], direct)
				if nums[j] > direct {
					return jump(nums[:i], true)
				}
				direct++
			}
			return jump(nums[:i], false)
		}
	}

	directs := make([]int, n)
	for i := range nums {
		directs[i] = n - 1 - i // 每个点都可直达
	}
	for i := n - 1 - 1; i >= 0; i-- {
		// fmt.Println(nums, directs, i, nums[i], directs[i])
		if nums[i] >= directs[i] { // 某个点能直达，找下一个点能否直达
			return jump(nums[:i+1], true)
		}
		return false
	}
	return true
}
