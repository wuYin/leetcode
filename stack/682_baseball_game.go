package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))                 // 30
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"})) // 27
}

// 有点后缀表达式计算值的感觉
func calPoints(ops []string) int {
	var s Stack
	for _, str := range ops {
		switch str {
		case "+":
			if !s.IsEmpty() {
				originPeek := s.Pop()
				sum := originPeek + s.Peek() // 所有 case 前两个元素都已是数字
				s.Push(originPeek)
				s.Push(sum)
			}
		case "D":
			if !s.IsEmpty() {
				s.Push(2 * s.Peek())
			}
		case "C":
			if !s.IsEmpty() {
				s.Pop()
			}
		default:
			n, _ := strconv.Atoi(str)
			s.Push(n)
		}
	}

	sum := 0
	for !s.IsEmpty() {
		sum += s.Pop()
	}

	return sum
}
