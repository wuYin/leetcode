package main

import "fmt"

// 定义二叉树结构
type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 按数组顺序插入节点
func (tree *Tree) BFSInsert(v int) {
	if tree.root == nil {
		tree.root = &TreeNode{Val: v}
		return
	}

	q := make([]*TreeNode, 0)
	q = append(q, tree.root)
	for len(q) > 0 {
		curNode := q[0] // 非递归分层遍历
		q = q[1:]       // pop

		if curNode.Left != nil { // 左子节点
			q = append(q, curNode.Left)
		} else {
			curNode.Left = &TreeNode{Val: v} // 找到合适的插入地点
			return
		}

		if curNode.Right != nil {
			q = append(q, curNode.Right)
		} else {
			curNode.Right = &TreeNode{Val: v}
			return
		}
	}
}

// 按层遍历
func (tree *Tree) BFSTraverse(node *TreeNode) {
	if node == nil {
		return
	}
	q := []*TreeNode{tree.root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		fmt.Printf("%v ", cur.Val)
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
}
