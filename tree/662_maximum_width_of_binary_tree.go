package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(widthOfBinaryTree(root)) // 2
}

// 满二叉树的节点宽度性质
//         1
//    2          3
//  4   5      6   7
// 8 9 10 11 12 13 14 15

// 理解宽度：最右边与最左边的数值之差
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var width int
	q := []*TreeNode{root}
	for len(q) > 0 {
		counts := len(q)
		width = max(width, q[len(q)-1].Val-q[0].Val+1) // 宽度为差值加 1
		for i := 0; i < counts; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				cur.Left.Val = cur.Val * 2
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				cur.Right.Val = cur.Val*2 + 1
				q = append(q, cur.Right)
			}
		}
	}
	return width
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
