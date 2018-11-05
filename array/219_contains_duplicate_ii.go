package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)) // true
}

// 查找还是哈希表比较 ok
func containsNearbyDuplicate(nums []int, k int) bool {
	duplicates := make(map[int][]int)
	for i, num := range nums {
		duplicates[num] = append(duplicates[num], i) // 把索引存起来
	}

	for _, dups := range duplicates {
		if len(dups) <= 1 {
			continue
		}
		for i := len(dups) - 1; i > 0; i-- {
			if dups[i]-dups[i-1] <= k { // 计算彼此的间隔
				return true
			}
		}
	}
	return false
}
