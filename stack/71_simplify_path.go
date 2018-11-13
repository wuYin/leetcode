package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/.."))
}

// package filepath 了解一下
func simplifyPath(path string) string {
	var s StrStack
	paths := strings.Split(path, "/")
	for _, p := range paths {
		if len(p) == 0 || p == "." {
			continue
		}
		if p == ".." {
			if !s.IsEmpty() { // 注意处理case: "/.."
				s.Pop()
			}
		} else {
			s.Push(p)
		}
	}
	return "/" + strings.Join(s, "/")
}
