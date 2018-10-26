package main

import "fmt"

func main() {
	fmt.Println(isValid("["))
	fmt.Println(isValid("([)]"))
}

// 类似后缀表达式的匹配过程
func isValid(s string) bool {
	var runes []rune
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			runes = append(runes, r)
		case ')', ']', '}':
			if len(runes) > 0 && opposite(r) == runes[len(runes)-1] {
				runes = runes[:len(runes)-1]
			} else {
				return false
			}
		}
	}
	if len(runes) > 0 {
		return false
	}
	return true
}

func opposite(r rune) rune {
	switch r {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	}
	return '\n'
}
