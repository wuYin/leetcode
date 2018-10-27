package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("abcd", "abcdabcdabcd")) // 3
	fmt.Println(repeatedStringMatch("abc", "cabcabca"))      // 4
}

// b 最多要 a 重复 a+Na+a 次
func repeatedStringMatch(A string, B string) int {
	count := 1
	a := A
	maxCount := len(B)/len(A) + 2
	for count <= maxCount {
		if strings.Contains(A, B) {
			return count
		}
		A += a
		count++
	}
	return -1
}
