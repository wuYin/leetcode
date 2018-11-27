package main

import "fmt"

func main() {
	post := []int{1, 5, 6, 4, 8, 26, 13, 7}
	in := []int{1, 4, 6, 5, 7, 8, 13, 26}
	root := buildTree(in, post)
	fmt.Println(root.Val)
	fmt.Println(root.Left)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Right)
	fmt.Println(root.Right.Right)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder)
}

func build(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}

	// 后序遍历的最后一个节点为根节点
	rootV := postorder[len(postorder)-1]
	var i int
	for ; i < len(inorder); i++ {
		if inorder[i] == rootV {
			break
		}
	}
	root := &TreeNode{Val: rootV}
	root.Left = build(inorder[:i], postorder[:i])                    // inorder 向前为左子树
	root.Right = build(inorder[i+1:], postorder[i:len(postorder)-1]) // inorder 向后为右子树
	return root
}
