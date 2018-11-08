package main

import "fmt"

func main() {
	fmt.Printf("%q\n", summaryRanges([]int{0, 1, 2, 4, 5, 7})) // ["0->2" "4->5" "7"]
}

// 要找到被剔除的数或数序列，判断某个数是否在数组中，自然想到哈希表
func summaryRanges(nums []int) []string {
	m := make(map[int]int)
	for _, n := range nums {
		m[n] = 0
	}

	var ranges []string
	l, r := -1, -1
	for _, num := range nums {
		_, preOK := m[num-1]
		_, nextOK := m[num+1]

		switch {
		case !preOK && nextOK: // num 在左边界
			l = num
		case preOK && !nextOK: // num 在右边界
			r = num
			ranges = append(ranges, fmt.Sprintf("%d->%d", l, r))
		case !preOK && !nextOK: // num 是独立的数
			ranges = append(ranges, fmt.Sprintf("%d", num))
		}
	}
	return ranges
}
