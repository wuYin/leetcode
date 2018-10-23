package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("0P"))
}

// 辣鸡解法
// TODO: 优化效率
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	str := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			str += string(r)
		}
	}
	head, tail := 0, len(str)-1
	for head < tail {
		if str[head] != str[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}
