package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	debug(deleteNode(root, 5)) // 2 3 4 6 7  // ok
}

func debug(root *TreeNode) {
	if root == nil {
		return
	}
	debug(root.Left)
	fmt.Print(root.Val, " ")
	debug(root.Right)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	var parent, cur *TreeNode = nil, root

	// 查找
	for cur != nil && cur.Val != key {
		parent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	switch {
	case cur == nil: // 没找到
		return root
	case parent == nil: // 需要去掉根节点，则将左子树嫁接到右子树上
		return insertNode(root.Right, root.Left)
	case cur.Val < parent.Val:
		parent.Left = nil
	case cur.Val > parent.Val:
		parent.Right = nil
	}

	// 将左子树和右子树转移到父节点上
	if cur.Left != nil {
		insertNode(root, cur.Left)
	}
	if cur.Right != nil {
		insertNode(root, cur.Right)
	}
	return root
}

func insertNode(root, subNode *TreeNode) *TreeNode {
	if root == nil {
		return subNode
	}
	if subNode == nil {
		return root
	}
	var parent, cur *TreeNode = nil, root
	for cur != nil {
		parent = cur
		if cur.Val > subNode.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	if parent.Val < subNode.Val {
		parent.Right = subNode
	} else {
		parent.Left = subNode
	}
	return root
}
