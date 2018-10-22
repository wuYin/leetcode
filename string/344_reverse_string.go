package main

import "fmt"

func main() {
	fmt.Println(reverseString("hello"))
}

// 字符串可索引，但不可原地更改值
func reverseString(s string) string {
	n := len(s)
	runes := make([]rune, 0, n)
	for i := n - 1; i >= 0; i-- {
		runes = append(runes, rune(s[i]))
	}
	return string(runes)
}
