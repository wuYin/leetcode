package main

import "fmt"

func main() {
	pre := []int{7, 4, 1, 6, 5, 13, 8, 26}
	in := []int{1, 4, 6, 5, 7, 8, 13, 26}
	root := buildTree(pre, in)
	fmt.Println(root.Val)
	fmt.Println(root.Left)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Left.Right)
	fmt.Println("----------")
	fmt.Println(root.Right)
	fmt.Println(root.Right.Left)
	fmt.Println(root.Right.Right)
}

// 还原二叉树
// 根据先序遍历或后序遍历，先找出根节点，确定好左子树与右子树
// 对左子树进行如上递归分析，对右子树进行递归分析
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder)
}

func build(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	root := &TreeNode{Val: preorder[0]}
	var i int
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	lPreorder := preorder[1 : i+1]
	rPreorder := preorder[i+1:]
	root.Left = build(lPreorder, inorder[:i])
	root.Right = build(rPreorder, inorder[i+1:])
	return root
}
