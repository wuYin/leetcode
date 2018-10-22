package main

import "fmt"

func main() {
	fmt.Println(toLowerCase("Hello"))
}

func toLowerCase(str string) string {
	const charDiff = 'a' - 'A'
	newStr := ""
	for _, r := range []rune(str) {
		if r < 'A' || r > 'Z' {
			newStr += string(r) // 非大写字母
			continue
		}
		newStr += string(r + charDiff)
	}
	return newStr
}
