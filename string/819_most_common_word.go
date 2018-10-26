package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}

// 没啥好说的，无脑哈希表
func mostCommonWord(paragraph string, banned []string) string {

	s := strings.ToLower(paragraph)
	words := strings.FieldsFunc(s, func(r rune) bool {
		if r >= 'a' && r <= 'z' {
			return false
		}
		return true
	})

	bannedM := make(map[string]interface{})
	for _, word := range banned {
		bannedM[word] = nil
	}
	wordsM := make(map[string]int)
	for _, word := range words {
		wordsM[word]++
	}

	for word := range wordsM {
		if _, ok := bannedM[word]; ok {
			delete(wordsM, word)
		}
	}

	maxCounts, mostCommon := 0, ""
	for word, count := range wordsM {
		if count > maxCounts {
			mostCommon = word
			maxCounts = count
		}
	}
	return mostCommon
}
