package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	flatten(root)

	fmt.Println(root)
	fmt.Println(root.Right)
	fmt.Println(root.Right.Right)
	fmt.Println(root.Right.Right.Right)
	fmt.Println(root.Right.Right.Right.Right)
}

// 原地展开，不能先序遍历建新树
//    1
//   / \
//  2   5

// 1
//  \
//   2
//    \
//     5
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	if root.Left != nil {
		lastNode := root.Right
		root.Right = root.Left // 把左子树挪到右子树上

		cur := root
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = lastNode // 把旧右子树接到新右子树
		root.Left = nil      // 清空左子树
	}
}
