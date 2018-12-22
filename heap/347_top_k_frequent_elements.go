package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{2, 2, 1, 1, 1, 3}, 2)) // [1, 2]
}

// 和 451 很像
// 自己重新定义优先队列的优先度排序即可
type Priority struct {
	n int // num
	p int // priority
}

func topKFrequent(nums []int, k int) []int {
	n2p := make(map[int]int)
	for _, n := range nums {
		n2p[n]++
	}

	var priorities []Priority
	for n, p := range n2p {
		priorities = append(priorities, Priority{n, p})
	}

	sort.Slice(priorities, func(i, j int) bool {
		return priorities[i].p > priorities[j].p
	})

	res := make([]int, k) // 应对 k 做合法校验
	for i, p := range priorities[:k] {
		res[i] = p.n
	}
	return res
}
