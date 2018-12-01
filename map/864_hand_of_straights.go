package main

import (
	"fmt"
	"sort"
)

// map first blood
func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3)) // true
}

func isNStraightHand(hand []int, W int) bool {
	m := make(map[int]int) // 计数 map
	for _, h := range hand {
		m[h]++
	}

	var ks []int
	for k := range m {
		ks = append(ks, k)
	}
	sort.Ints(ks)

	// 好戏开始
	for _, k := range ks {
		c1 := m[k]
		if c1 <= 0 {
			continue // 已经被完全拖下水
		}
		for i := k; i < k+W; i++ {
			if c2, ok := m[i]; !ok || c2 < c1 {
				return false
			}
			m[i] -= c1 // 一次走一批
		}
	}

	return true
}
