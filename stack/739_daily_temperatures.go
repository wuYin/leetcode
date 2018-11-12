package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

// 十分有栈味道的题
// 弄清楚栈在题中如何使用，灵活结合 pop push 等操作解决
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var s Stack // 索引栈
	for i, t := range temperatures {
		for !s.IsEmpty() {
			peek := s.Peek()
			if t > temperatures[peek] { // 遇到大数，则把积压的小数全部 pop 出来
				res[peek] = i - peek
				s.Pop()
			} else {
				break
			}
		}
		s.Push(i)
	}
	return res
}
