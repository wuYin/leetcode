package main

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 1}

	//    4
	//   / \
	//  9   0
	// / \
	// 5   1
	fmt.Println(preorderTraversal(root)) // [4 9 5 1 0]
}

func preorderTraversal(root *TreeNode) []int {
	var nums []int
	traverse(root, &nums)
	return nums
}

func traverse(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	traverse(root.Left, nums)
	traverse(root.Right, nums)
}

// TODO
// iteration
