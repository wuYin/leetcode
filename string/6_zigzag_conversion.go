package main

import "fmt"

func main() {
	fmt.Println(convert("ABCDEFGHIJKLMN", 4)) // AGMBFHLNCEIKDJ
}

func convert(s string, numRows int) string {
	n := len(s)
	if n <= 1 || numRows <= 1 {
		return s
	}

	lines := make([][]rune, numRows)
	for i := 0; i < n; {
		// 向下走
		for r := 0; r <= numRows-1 && i < n; r++ {
			lines[r] = append(lines[r], rune(s[i]))
			i++
		}
		// 向右上方走
		for r := numRows - 2; r >= 1 && i < n; r-- {
			lines[r] = append(lines[r], rune(s[i]))
			i++
		}
	}

	var str string
	for _, line := range lines {
		for _, r := range line {
			str += string(r)
		}
	}

	return str
}
