package main

import "fmt"

func main() {
	s := &TreeNode{Val: 3}
	s.Left = &TreeNode{Val: 4}
	s.Left.Left = &TreeNode{Val: 1}
	s.Left.Right = &TreeNode{Val: 2}
	// s.Left.Right.Left = &TreeNode{Val: 0} // false
	s.Right = &TreeNode{Val: 5}

	t := &TreeNode{Val: 4}
	t.Left = &TreeNode{Val: 1}
	t.Right = &TreeNode{Val: 2}

	fmt.Println(isSubtree(s, t)) // true
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSame(s, t) {
		return true // bingo
	}

	if s == nil {
		return false // 向下不会再相等
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
}
