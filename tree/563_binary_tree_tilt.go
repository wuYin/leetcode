package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(findTilt(root)) // 1
}

func findTilt(root *TreeNode) int {
	var titlSum int
	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lVal := traverse(node.Left)
		rVal := traverse(node.Right)
		titlSum += abs(lVal - rVal) // 需要所有差值，不断向上累计

		return node.Val + lVal + rVal // 返回整棵子树和
	}
	traverse(root)

	return titlSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
