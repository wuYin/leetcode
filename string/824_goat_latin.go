package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("Each word consists of lowercase and uppercase letters only"))
}

// 主要是捋清楚规则
func toGoatLatin(S string) string {
	words := strings.Fields(S)
	newWords := make([]string, 0, len(words))
	for i, word := range words {
		switch word[0] {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U': // 提及规则的时候，再简单的已知条件（元音：大小写）也要分析清楚
			word += "ma"
		default:
			word = word[1:] + string(word[0]) + "ma"
		}
		i += 1
		for i > 0 {
			word += "a"
			i--
		}
		newWords = append(newWords, word)
	}
	return strings.Join(newWords, " ")
}
