package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "ab"))
}

// 理清楚逻辑即可
func canConstruct(ransomNote string, magazine string) bool {
	m1, m2 := make(map[rune]int), make(map[rune]int)
	for _, r := range ransomNote {
		m1[r]++
	}
	for _, r := range magazine {
		m2[r]++
	}
	for r, c1 := range m1 {
		if c2, ok := m2[r]; ok {
			if c1 > c2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
