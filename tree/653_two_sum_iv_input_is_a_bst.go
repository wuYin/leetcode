package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	fmt.Println(findTarget(root, 2)) // false
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]int)
	return preTraverse(root, k, m)
}

// 和 two sum 一样
func preTraverse(root *TreeNode, k int, m map[int]int) bool {
	if root == nil {
		return false
	}

	if _, ok := m[k-root.Val]; ok { // 先序遍历保证了去重
		return true
	}

	m[root.Val]++
	return preTraverse(root.Left, k, m) || preTraverse(root.Right, k, m)
}
