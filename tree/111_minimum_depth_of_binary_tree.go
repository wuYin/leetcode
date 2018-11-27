package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 1}
	for _, v := range []int{2} {
		tree.BFSInsert(v) // 自己写的破 insert 调试半天  fuck
	}
	fmt.Println(minDepth(tree.root)) // 2
}

func minDepth(root *TreeNode) int {
	return depth(root)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := depth(root.Left)
	rDepth := depth(root.Right)

	if lDepth == 0 { // 左边走到尽头，只能往右边走
		return rDepth + 1
	}
	if rDepth == 0 { // 往左边走
		return lDepth + 1
	}

	return min(lDepth, rDepth) + 1 // 两边都可以走，选最浅的子树
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
