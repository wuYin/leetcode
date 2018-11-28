package main

import "fmt"

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}

	//   3
	//  / \
	// 1   4
	//  \
	//   2
	fmt.Println(kthSmallest(root, 1)) // 1
}

// 显然是中序遍历的 stack 迭代遍历，计数取值
func kthSmallest(root *TreeNode, k int) int {
	if k == 0 || root == nil {
		return -1
	}

	var counter int
	var stack []*TreeNode
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		// 现在左子节点为空 // bingo
		parent := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop
		counter++
		if counter == k {
			return parent.Val
		}
		cur = parent.Right
	}
	return -1
}
