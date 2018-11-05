package main

import "fmt"

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))  // 2
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, -1)) // 0
}

// 和 two sum 一样，用哈希表做统计
func findPairs(nums []int, k int) int {
	if k < 0 { // 注意这个case...
		return 0
	}

	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	}

	count := 0
	if k == 0 {
		for _, freq := range freqs {
			if freq > 1 {
				count++
			}
		}
		return count
	}

	for num := range freqs {
		if _, ok := freqs[num+k]; ok {
			count++
		}
	}
	return count
}
