package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(badFrequencySort("tree")) // eetr
	fmt.Println(frequencySort("tree"))    // eetr
}

// heap 数组实现的优先队列左右子节点的优先度不分大小
// 所以需要重新组织优先度的排序
type Counter struct {
	r rune // rune
	p int  // priority // 出现的次数
}

func frequencySort(s string) string {
	r2c := make(map[rune]int)
	for _, r := range s {
		r2c[r]++
	}

	var counters []Counter
	for r, c := range r2c {
		counters = append(counters, Counter{r, c})
	}

	// 自定义优先度排序规则
	sort.Slice(counters, func(i, j int) bool {
		return counters[i].p > counters[j].p
	})

	s = ""
	for _, counter := range counters {
		s += strings.Repeat(string(counter.r), counter.p)
	}
	return s
}

// 最直观的硬核实现
func badFrequencySort(s string) string {
	r2c := make(map[rune]int)
	for _, r := range s {
		r2c[r]++
	}

	c2rs := make(map[int][]rune)
	for r, c := range r2c {
		c2rs[c] = append(c2rs[c], r)
	}

	counts := make([]int, 0, len(c2rs))
	for c := range c2rs {
		counts = append(counts, c)
	}

	sort.Ints(counts)

	var res string
	for i := len(counts) - 1; i >= 0; i-- {
		for _, r := range c2rs[counts[i]] {
			res += strings.Repeat(string(r), counts[i])
		}
	}

	return res
}
