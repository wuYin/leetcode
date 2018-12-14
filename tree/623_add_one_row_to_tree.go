package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 6},
	}

	newRoot := addOneRow(root, 1, 2)
	fmt.Println(newRoot.Val, newRoot.Left.Val, newRoot.Right.Val) // 4 1 1
	// also ok
}

// 层序遍历即可
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return nil
	}
	if d == 1 {
		return &TreeNode{Val: v, Left: root} // 特殊 case
	}

	depth := 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		counts := len(q)
		for i := 0; i < counts; i++ {
			cur := q[0]
			q = q[1:]
			if depth == d-1 { // 注意 -1
				cur.Left = &TreeNode{Val: v, Left: cur.Left}
				cur.Right = &TreeNode{Val: v, Right: cur.Right}
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return root
}
