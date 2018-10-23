package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"a", "b", "c"}))                            // 3	// 自己和自己是特殊等价的
	fmt.Println(numSpecialEquivGroups([]string{"abc", "acb", "bac", "bca", "cab", "cba"})) // 3
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "adcb", "cbad"}))           // 1
}

// 不要暴力的用 map 和 inSlice 的多重遍历去判断，搞复杂了
// 抓住题目本质：特殊等价字符串，即奇数位、偶数位字符数组随意排列组合后仍认为是等价的
// 相同元素的不同排列组合，排序后是一致的。判断 2 个 []string 是否一致也是这个道理
func numSpecialEquivGroups(A []string) int {
	m := make(map[string]int, len(A)/2)
	// 所有
	for _, s := range A {
		var evens, odds []string
		for i := 0; i < len(s); i += 2 {
			evens = append(evens, string(s[i]))
		}
		for i := 1; i < len(s); i += 2 {
			odds = append(odds, string(s[i]))
		}
		sort.Strings(evens)
		sort.Strings(odds)
		sortedS := strings.Join(evens, "") + strings.Join(odds, "")
		m[sortedS]++
	}
	return len(m)
}
