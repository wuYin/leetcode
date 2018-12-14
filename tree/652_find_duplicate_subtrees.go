package main

import (
	"fmt"
	"strconv"
)

func main() {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Left = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 3}
	tree.Right.Left = &TreeNode{Val: 2}
	tree.Right.Left.Left = &TreeNode{Val: 4}
	tree.Right.Right = &TreeNode{Val: 4}

	for _, node := range findDuplicateSubtrees(tree) {
		fmt.Println(node.Val) // ok
	}
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	paths := make(map[string]int)
	var res []*TreeNode
	traverse(root, paths, &res)
	return res
}

// 后序遍历把路径放到 map 中去重
func traverse(root *TreeNode, paths map[string]int, res *[]*TreeNode, ) string {
	if root == nil {
		return "#"
	}

	fullPath := strconv.Itoa(root.Val) + "," + traverse(root.Left, paths, res) + "," + traverse(root.Right, paths, res)
	if count, ok := paths[fullPath]; ok && count == 1 { // 计数 count 为 1
		*res = append(*res, root)
	}
	paths[fullPath]++

	return fullPath
}
