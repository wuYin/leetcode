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
	fmt.Println(postorderTraversal(root)) // [5 1 9 0 4]
}

func postorderTraversal(root *TreeNode) []int {
	var nums []int
	traverse(root, &nums)
	return nums
}

func traverse(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, nums)
	traverse(root.Right, nums)
	*nums = append(*nums, root.Val)
}

// TODO
// iteration
