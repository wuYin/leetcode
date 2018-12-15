package main

import "fmt"

func main() {
	t1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	t2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(flipEquiv(t1, t2)) // true
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}
