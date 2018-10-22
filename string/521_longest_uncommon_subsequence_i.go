package main

import "fmt"

func main() {
	fmt.Println(findLUSlength("aba", "cdc"))
}

// 哈希表长度不大，两两比较是很好的选择
func findLUSlength(a string, b string) int {
	m1, m2 := subStr(a), subStr(b)
	maxSubSeq := 0
	for s := range m1 {
		if _, ok := m2[s]; !ok && len(s) > maxSubSeq {
			maxSubSeq = len(s)
		}
	}
	for s := range m2 {
		if _, ok := m1[s]; !ok && len(s) > maxSubSeq {
			maxSubSeq = len(s)
		}
	}
	if maxSubSeq == 0 {
		return -1
	}
	return maxSubSeq
}

func subStr(s string) map[string]int {
	n := len(s)
	m := make(map[string]int, (1+n)*n/2) // 等差数列长度
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			m[s[i:j+1]]++
		}
	}
	return m
}
