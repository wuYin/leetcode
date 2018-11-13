package main

// 元素类型不一致，蛋疼
// 可以试试 genny 类型生成器
type StrStack []string

func (s *StrStack) Len() int {
	return len(*s)
}

func (s *StrStack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *StrStack) Push(elem string) {
	*s = append(*s, elem)
}

func (s *StrStack) Pop() string {
	n := s.Len()
	peek := (*s)[n-1]
	*s = (*s)[:n-1]
	return peek
}

func (s *StrStack) Peek() string {
	n := s.Len()
	return (*s)[n-1]
}
