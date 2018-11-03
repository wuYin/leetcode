package main

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("aaa"))               // [[0,2]]
	fmt.Println(largeGroupPositions("abbxxxxzzy"))        // [[3,6]]
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd")) // [[3,5] [6,9] [12,14]]
}

func largeGroupPositions(S string) [][]int {
	n := len(S)
	if n <= 2 {
		return nil
	}

	var locs [][]int
	slow, fast := 0, 0
	for slow < n {
		for fast < n && S[fast] == S[slow] {
			fast++ // 注意跳出循环时 S[fast] != S[slow]
		}
		if fast-slow >= 3 {
			locs = append(locs, []int{slow, fast - 1})
		}

		if fast > slow {
			slow = fast - 1 // 抵消下边一般 case 的 +1
		}
		slow++
		fast = slow
	}
	return locs
}
