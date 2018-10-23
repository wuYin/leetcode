package main

import "fmt"

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("Leetcode"))
	fmt.Println(detectCapitalUse("mL"))
}

// 训练逻辑, it's ok
func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}

	firstUp := rune(word[0]) >= 'A' && rune(word[0]) <= 'Z'
	allUpCase, allLowCase := true, true
	for i := 1; i < len(word); i++ {
		r := rune(word[i])
		if r < 'A' || r > 'Z' {
			allUpCase = false
		}
		if r < 'a' || r > 'z' {
			allLowCase = false
		}
	}

	if firstUp && (allUpCase || allLowCase) {
		return true
	}
	if !firstUp && allLowCase {
		return true
	}
	return false
}
