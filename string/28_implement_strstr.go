package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("a", ""))
}

func strStr(haystack string, needle string) int {
	if len(needle) <= 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	hs, ns := []rune(haystack), []rune(needle)
	lh, ln := len(hs), len(ns)
	var starts []int
	for i, h := range hs {
		if h == ns[0] {
			starts = append(starts, i)
		}
	}
	for _, start := range starts {
		if start+ln > lh {
			break
		}
		if string(hs[start:start+ln]) == needle {
			return start
		}
	}
	return -1
}
