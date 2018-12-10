package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 13}
	newRoot := convertBST(root)
	fmt.Println(newRoot.Val)       // 18
	fmt.Println(newRoot.Left.Val)  // 20
	fmt.Println(newRoot.Right.Val) // 13
}

func convertBST(root *TreeNode) *TreeNode {
	var curSum int
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Right) // 也是后序遍历，不过方向改了
		curSum += node.Val
		node.Val = curSum
		traverse(node.Left)
	}

	traverse(root)
	return root
}
