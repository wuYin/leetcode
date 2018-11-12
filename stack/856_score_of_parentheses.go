package main

import "fmt"

func main() {
	fmt.Println(scoreOfParentheses("(())")) // 2
}

// 可以，这很栈
// 简化问题，解决 (()) 和 (()()) 即可
func scoreOfParentheses(S string) int {
	var s Stack
	for _, r := range S {
		switch r {
		case '(':
			s.Push(-1)
		case ')':
			cur := 0
			for s.Peek() != -1 {
				cur += s.Pop()
			}
			s.Pop()
			if cur == 0 {
				s.Push(1)
			} else {
				s.Push(cur * 2)
			}
		}
	}

	sum := 0
	for !s.IsEmpty() {
		sum += s.Pop()
	}
	return sum
}
