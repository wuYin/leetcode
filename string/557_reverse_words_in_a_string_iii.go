package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

// Field Split Join 是好搭档
func reverseWords(s string) string {
	reverse := func(word string) string {
		n := len(word)
		runes := make([]rune, 0, n)
		for i := n - 1; i >= 0; i-- {
			runes = append(runes, rune(word[i]))
		}
		return string(runes)
	}
	words := strings.Fields(s)
	res := make([]string, 0, len(words))
	for _, word := range words {
		res = append(res, reverse(word))
	}
	return strings.Join(res, " ")
}
