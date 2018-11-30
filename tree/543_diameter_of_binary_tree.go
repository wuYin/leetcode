package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(diameterOfBinaryTree(root))
}

// 显然是找左子树最大路径与右子树最长路径
func diameterOfBinaryTree(root *TreeNode) int {
	var d int
	traverse(root, &d)
	return d
}

// 递归解法比较辣鸡
func traverse(root *TreeNode, maxD *int) {
	if root == nil {
		return
	}
	curD := depth(root.Left) + depth(root.Right)
	if curD > *maxD {
		*maxD = curD
	}

	traverse(root.Left, maxD)
	traverse(root.Right, maxD)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := depth(root.Left) + 1
	rDepth := depth(root.Right) + 1
	if lDepth > rDepth {
		return lDepth
	}
	return rDepth
}
