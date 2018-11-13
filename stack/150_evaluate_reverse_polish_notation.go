package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"4", "3", "-"}))                                                       // 1
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 22
}

// 再加上 () 运算就是完整的表达式计算了
func evalRPN(tokens []string) int {
	var s Stack
	for _, token := range tokens {
		if n, err := strconv.Atoi(token); err == nil {
			s.Push(n)
		} else {
			// 字符运算，注意减除的顺序
			l, r := s.Pop(), s.Pop()
			switch token {
			case "+":
				s.Push(l + r)
			case "-":
				s.Push(r - l)
			case "*":
				s.Push(l * r)
			case "/":
				s.Push(r / l)
			}
		}
	}
	return s.Pop()
}
