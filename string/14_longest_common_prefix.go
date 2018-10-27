package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "fly", "flight"})) // fl
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))   // ""
}

// 维护一个前缀库，不断往后遍历，判断前缀逐步剪短前缀库的大小
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	s1 := strs[0]
	prefixes := make([]string, len(s1))
	for i := range s1 {
		prefixes[i] = s1[:i+1]
	}

	for _, s := range strs {
		for i, pre := range prefixes {
			if !strings.HasPrefix(s, pre) {
				prefixes = prefixes[:i]
				break
			}
		}
	}

	if len(prefixes) > 0 {
		return prefixes[len(prefixes)-1]
	}
	return ""
}
