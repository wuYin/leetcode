package main

func main() {
	// nextPermutation([]int{1, 2, 3})
	// nextPermutation([]int{3, 2, 1})
	// nextPermutation([]int{1, 1, 5})
	// nextPermutation([]int{1, 2, 7, 4, 3, 1})
	nextPermutation([]int{1, 3, 3})
}

//
// 题目的本质是一般规律：从后向前找出第一个降序点，从前向后找向上最接近的数，调换，将后面序列转置保证后序是升序为最小排列
// 遇到这类题，使用长数组做示例，再找其中的规律。不要硬磕用位数去解决，看清楚题目的本质后用尽可能简单的方式解决，并满足复杂度要求
//
func nextPermutation(nums []int) {
	n := len(nums)
	// 先检查是否为最大序
	isMaxSeq := true
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			isMaxSeq = false
		}
	}
	// 已是最大排列，则原地旋转
	if isMaxSeq {
		reverse(nums)
		return
	}
	// 向前遍历找出要调换的值
	pre := n - 1
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pre = i - 1 // 找到第一个降序点 // 2
			break
		}
	}
	nearest := pre + 1
	for i := pre + 1; i < n; i++ {
		if nums[i] > nums[pre] && nums[i] <= nums[nearest] { // <= 找到最靠后的向上最接近的数 // 3
			nearest = i
		}
	}
	// 交换二者的位置	// 1 3 7 4 2 1
	nums[pre], nums[nearest] = nums[nearest], nums[pre]
	// 倒置 pre+1 之后的数据，保证后续数据是升序	// 1 3 1 2 4 7
	reverse(nums[pre+1:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
