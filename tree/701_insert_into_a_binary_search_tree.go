package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	insertIntoBST(root, 2)
	fmt.Println(root)
}

// 递归赋值
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	switch {
	case root.Val > val:
		root.Left = insertIntoBST(root.Left, val)
	case root.Val < val:
		root.Right = insertIntoBST(root.Right, val)
	default:
	}
	return root // 相同值不予处理
}
