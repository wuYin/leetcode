package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(tree2str(root)) // 1(2()(4))(3)
}

func tree2str(t *TreeNode) string {
	s := traverse(t)
	if len(s) > 2 {
		s = s[1 : len(s)-1]
	}
	return s
}

func traverse(root *TreeNode) string {
	if root == nil {
		return ""
	}

	l := traverse(root.Left)
	r := traverse(root.Right)

	if r != "" && l == "" { // 要左不要右
		l = "()"
	}
	return fmt.Sprintf("(%d%s%s)", root.Val, l, r)
}
