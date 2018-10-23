package main

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
}

// 理解题意：相同数量 0 和 1 的子串
// 统计遍历考虑处理第一个元素后，从第二个开始向前遍历
// 相比于从第一个遍历到倒数第二个要简洁易处理
func countBinarySubstrings(s string) int {
	if len(s) <= 0 {
		return 0
	}

	zeroCount, oneCount := 0, 0
	if s[0] == '0' {
		zeroCount++
	} else {
		oneCount++
	}

	counts := 0
	for i := 1; i < len(s); i++ {
		switch s[i] {
		case '0':
			if s[i-1] == '0' {
				zeroCount++
			} else {
				zeroCount = 1
			}
			if oneCount >= zeroCount { // 对称
				counts++
			}
		case '1':
			if s[i-1] == '1' {
				oneCount++
			} else {
				oneCount = 1
			}
			if zeroCount >= oneCount {
				counts++
			}
		}
	}
	return counts
}
