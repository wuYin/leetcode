package main

import "fmt"

func main() {
	root := sortedArrayToBST([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Left.Right)
	fmt.Println(root.Right)
	fmt.Println(root.Right.Left) // bingo
}

func sortedArrayToBST(nums []int) *TreeNode {
	return construct(nums)
}

func construct(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	// 尽可能地取中间的值作为根节点
	mid := n / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = construct(nums[:mid])    // 左子数组作为左子树
	root.Right = construct(nums[mid+1:]) // 右子数组作为右子树 // 地达到 AVL 树特点
	return root
}
