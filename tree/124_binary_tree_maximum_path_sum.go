package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	// root.Right.Left = &TreeNode{Val: 15}
	// root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(maxPathSum(root))
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := -1 << 31
	pathSum(root, &max)
	return max
}

func pathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	lSum := pathSum(node.Left, maxSum)
	rSum := pathSum(node.Right, maxSum)

	curSum := node.Val
	if lSum > 0 {
		curSum += lSum
	}
	if rSum > 0 {
		curSum += rSum
	}

	*maxSum = max(*maxSum, curSum)
	return max(node.Val, node.Val+lSum, node.Val+rSum)
}

func max(nums ...int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}
