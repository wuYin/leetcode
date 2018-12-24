package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(myAtoi("-10"))                 // -10
	fmt.Println(myAtoi("9223372036854775808")) // 2147483647
}

const (
	INT_MIN = -1 << 31
	INT_MAX = 1<<31 - 1
)

// 正常的整数字符串只有 3 种case
// 10
// +10
// -10
// 注意处理溢出的情形
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	positive := true
	switch str[0] {
	case '-':
		positive = false
		str = str[1:]
	case '+':
		str = str[1:]
	}

	var nums []rune
	for _, r := range str {
		if r < '0' || r > '9' {
			break
		}
		nums = append(nums, r)
	}

	var n int
	for _, num := range nums {
		n = 10*n + int(num-'0')
		if positive && n > INT_MAX {
			return INT_MAX
		}
		if !positive && -n < INT_MIN {
			return INT_MIN
		}
	}

	if !positive {
		return -n
	}
	return n
}
