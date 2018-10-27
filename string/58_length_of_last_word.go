package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

// 捡了个 Golang 标准库的小便宜...
func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	if len(words) <= 0 {
		return 0
	}
	return len(words[len(words)-1])
}
