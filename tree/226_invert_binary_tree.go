package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 4}
	for _, v := range []int{2, 7, 1, 3, 6, 9} {
		tree.BFSInsert(v)
	}
	tree.BFSTraverse(tree.root) // 4 2 7 1 3 6 9
	fmt.Println()
	invertTree(tree.root)
	tree.BFSTraverse(tree.root) // 4 7 2 9 6 3 1
}

// 反转二叉树其实是镜像翻转 swap 节点操作
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)

	swapNode := root.Left // 镜像翻转
	root.Left = root.Right
	root.Right = swapNode
	return root
}
