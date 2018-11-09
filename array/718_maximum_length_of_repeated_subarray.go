package main

import "fmt"

func main() {
	fmt.Println(findLength([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}))                               // 2
	fmt.Println(findLength([]int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0})) // 9
}

// 最直观也是最辣鸡的做法
func findLength(A []int, B []int) int {
	m := make(map[int][]int)
	for i, b := range B {
		m[b] = append(m[b], i) // 存储 B 数组中所有元素的索引
	}

	maxLen := 0
	for i, a := range A {
		if _, ok := m[a]; !ok { // 不存在与 B 数组中
			continue
		}
		// 向后遍历
		for _, start := range m[a] {
			curLen := 0
			ai := i
			for bi := start; ai < len(A) && bi < len(B); ai, bi = ai+1, bi+1 { // 连续相同的子数组计数
				if A[ai] != B[bi] {
					break
				}
				curLen++
			}
			if maxLen < curLen {
				maxLen = curLen
			}
		}
	}

	return maxLen
}
