package main

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 1}
	fmt.Println(sumNumbers(root)) // 1026
}

// 和 112 113 一模一样
func sumNumbers(root *TreeNode) int {
	var sums []int
	traverse(root, 0, &sums)
	var sum int
	for _, v := range sums {
		sum += v
	}
	return sum
}

// 遍历求和并取叶子节点即可
func traverse(root *TreeNode, curSum int, nums *[]int) {
	if root == nil {
		return
	}

	curSum *= 10
	curSum += root.Val
	if root.Left == nil && root.Right == nil {
		*nums = append(*nums, curSum)
	}

	traverse(root.Left, curSum, nums)
	traverse(root.Right, curSum, nums)
}
