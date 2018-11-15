package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 3}
	for _, v := range []int{9, 20, -1, -1, 15, 7} {
		tree.BFSInsert(v)
	}
	fmt.Println(maxDepth(tree.root)) // 3
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}
