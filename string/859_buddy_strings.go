package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("", ""))         // false
	fmt.Println(buddyStrings("aa", "aa"))     // true
	fmt.Println(buddyStrings("ab", "ab"))     // false
	fmt.Println(buddyStrings("abab", "abab")) // true
	fmt.Println(buddyStrings("abcd", "abcd")) // false
}


// 捋清楚逻辑，找规律...
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) || A == "" || B == "" {
		return false
	}
	diff := make([]int, 0, 2)
	for index := range B {
		if A[index] != B[index] {
			diff = append(diff, index)
		}
	}
	switch len(diff) {
	case 0:
		repeat := false
		for i := 1; i < len(A); i++ {
			if A[0] == A[i] {
				repeat = true
			}
		}
		return repeat
	case 1:
		return false
	case 2:
		runes := []rune(A)
		runes[diff[0]], runes[diff[1]] = runes[diff[1]], runes[diff[0]]
		return string(runes) == B
	}
	return false
}
