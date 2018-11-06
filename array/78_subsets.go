package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}

// 自创的 4ms 解法阿西吧
// 理解子集的本质，N 个数有 1~N 种组合方式
// 可以直接从后向前遍历，每遇到一个数就把它与所有的现有子集重新 append 成新子集
// 遍历完毕就是所有子集
// 我觉得 ok
func subsets(nums []int) [][]int {
	n := len(nums)
	seqs := make([][]int, 1)
	seqs = append(seqs, []int{nums[n-1]})

	for i := n - 2; i >= 0; i-- {
		for _, seq := range seqs {
			nextSeq := append([]int{nums[i]}, seq...)
			seqs = append(seqs, nextSeq)
		}
	}

	return seqs
}
