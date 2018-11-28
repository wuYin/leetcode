package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}

	//    1
	//   / \
	//  2   3
	// / \  /
	// 4  5 6
	fmt.Println(countNodes(root)) // 6
}

func countNodes(root *TreeNode) int {
	var count int
	countNode(root, &count)
	return count
}

// 二叉树计数节点个数
// 应该归为 easy tag...
func countNode(node *TreeNode, count *int) {
	if node == nil {
		return
	}
	*count++
	countNode(node.Left, count)
	countNode(node.Right, count)
}
