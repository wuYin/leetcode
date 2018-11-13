package main

import "fmt"

func main() {
	fmt.Println(decodeString("3[a]2[b4[F]c]")) // aaabFFFFcbFFFFc
}

// 考虑 case
// 字母大小写
// 数字和字母不止一个，需进行进制和字符的暂存
// 分数字栈、字符栈进行计算，遇到 ] 触发一次计算

// go run stack.go str_stack.go 394_xxx.go
func decodeString(s string) string {
	var (
		n    int
		res  string
		nums Stack
		str  StrStack
	)
	for _, r := range s {
		switch {
		case '0' <= r && r <= '9':
			n = 10*n + int(r-'0')
		case ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z'):
			res += string(r)
		case r == '[': // 双双进栈
			nums.Push(n)
			n = 0
			str.Push(res)
			res = ""
		case r == ']': // 双双出栈计算值
			repeat := res
			for i := 0; i < nums.Peek()-1; i++ {
				res += repeat
			}
			res = str.Pop() + res
			nums.Pop()
		}
	}
	return res
}
