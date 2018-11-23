package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 4}
	for _, v := range []int{2, 7, 1, 3, 6, 9} {
		tree.BFSInsert(v)
	}
	fmt.Println(isValidBST(tree.root)) // true
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return verify(root, -1<<63, 1<<63-1)
}

// min 和 max 一路存储已验证的 bst 最小值和最大值
func verify(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val >= max || node.Val <= min {
		return false
	}
	return verify(node.Left, min, node.Val) && verify(node.Right, node.Val, max)
}
