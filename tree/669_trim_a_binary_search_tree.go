package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(trimBST(root, 1, 2)) // ok
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	// 检查 root 需要 trim 的情况
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
