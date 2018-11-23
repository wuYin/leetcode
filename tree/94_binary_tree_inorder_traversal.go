package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 4}
	for _, v := range []int{2, 7, 1, 3, 6, 9} {
		tree.BFSInsert(v)
	}
	tree.BFSTraverse(tree.root)              // 4 2 7 1 3 6 9
	fmt.Println(inorderTraversal(tree.root)) // 1 2 3 4 6 7 9
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	traverse(root, &res)
	return res
}

// 中序遍历：左子节点 -> 父节点 -> 右子节点
func traverse(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		traverse(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		traverse(node.Right, res)
	}
}
