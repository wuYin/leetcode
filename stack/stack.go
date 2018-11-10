package main

import (
	"errors"
	"log"
)

// 简单的栈定义
// go run stack.go n_xxx.go
type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(elem int) {
	*s = append(*s, elem)
}

var ERR_STACK_EMPTY = errors.New("stack is empty")

func (s *Stack) Pop() int {
	n := s.Len()
	if n == 0 {
		log.Fatal(ERR_STACK_EMPTY)
	}
	peek := (*s)[n-1]
	*s = (*s)[:n-1]
	return peek
}

func (s *Stack) Peek() int {
	n := s.Len()
	if n == 0 {
		log.Fatal(ERR_STACK_EMPTY)
	}
	return (*s)[n-1]
}
