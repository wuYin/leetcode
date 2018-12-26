// 参考 https://goleetcode.io/2018/11/20/array/31-next-permutation
package main

import "fmt"

func main() {
	nums := []int{1, 2, 7, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums) // [1 3 1 2 4 7] // bingo
}

// 数组规律题
// 从后往前找第一个下降点 i，再从后往前找它的 ceil 值，交换
// 再将 [i+1:] 之后的数据从降序反转为升序，为最小序列
func nextPermutation(nums []int) {
	// 处理降序的 case
	desc := true
	n := len(nums)
	for i := range nums[:n-1] {
		if nums[i] < nums[i+1] {
			desc = false
		}
	}
	if desc {
		reverse(nums)
		return
	}

	// 从后向前找第一个下降的点
	var i int
	for i = n - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			i-- // 找到 2
			break
		}
	}

	// 从后向前，找向上最接近的值
	for j := n - 1; j > i; j-- {
		if nums[j] > nums[i] {
			nums[j], nums[i] = nums[i], nums[j] // 交换 2 和 3 	// [1 3 7 4 2 1]
			break
		}
	}

	reverse(nums[i+1:]) // 反转 4 2 1	// [1 3 1 2 4 7]
}
