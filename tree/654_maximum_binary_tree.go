package main

import "fmt"

func main() {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(root) // ... //ok
}

// 直接取值恢复
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := 0
	max := nums[0]
	for i, n := range nums {
		if n > max {
			max = n
			mid = i
		}
	}

	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTree(nums[:mid])
	root.Right = constructMaximumBinaryTree(nums[mid+1:])
	return root
}
