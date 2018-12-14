package main

import "fmt"

func main() {
	t1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 2},
	}
	t2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}},
	}
	t3 := mergeTrees(t1, t2)
	fmt.Println(t3.Val) // 3
	// also ok
}

// DP
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}
}
