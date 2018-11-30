package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 0}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 1}
	newRoot := pruneTree(root)
	fmt.Println(newRoot)
	fmt.Println(newRoot.Left)
	fmt.Println(newRoot.Right)
	fmt.Println(newRoot.Right.Left)
	fmt.Println(newRoot.Right.Right) // bingo
}

// 其实就是一个后序遍历，一步一步摘除 0  叶子
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
