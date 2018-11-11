package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(backspaceCompare("a#c", "b")) // false
}

func backspaceCompare(S string, T string) bool {
	var s1, s2 Stack
	traverse := func(str string, s *Stack) {
		for _, r := range str {
			if r == '#' {
				if !s.IsEmpty() {
					s.Pop()
				}
			} else {
				s.Push(int(r))
			}
		}
	}

	traverse(S, &s1)
	traverse(T, &s2)

	var s, t string
	for !s1.IsEmpty() {
		s += string(rune(s1.Pop()))
	}
	for !s2.IsEmpty() {
		t += string(rune(s2.Pop()))
	}

	return strings.EqualFold(s, t)
}
