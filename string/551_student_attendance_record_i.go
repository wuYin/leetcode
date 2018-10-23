package main

import "fmt"

func main() {
	fmt.Println(checkRecord("PPALLL"))
}

// 2 个逻辑对应 2 个代码块
func checkRecord(s string) bool {
	countA := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			countA++
		}

		if i < len(s)-2 && s[i] == 'L' && s[i+1] == 'L' && s[i+2] == 'L' {
			return false
		}
	}
	return countA <= 1
}
