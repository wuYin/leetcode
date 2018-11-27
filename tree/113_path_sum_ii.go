package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 8}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Left = &TreeNode{Val: 5}

	// traverse 内部需要对 curPath 进行 copy 操作，否则结果就是 [[5 8 4 1]]
	fmt.Println(pathSum(root, 22))
}

func pathSum(root *TreeNode, sum int) [][]int {
	var paths [][]int
	traverse(root, 0, sum, nil, &paths)
	return paths
}

// 每个节点都只对应一种路径
// 这里 curPath 每个节点对应的值不同，不传指针
func traverse(root *TreeNode, curSum, target int, curPath []int, paths *[][]int) {
	if root == nil {
		return
	}

	curPath = append(curPath, root.Val)
	curSum = curSum + root.Val
	if root.Left == nil && root.Right == nil && curSum == target {
		*paths = append(*paths, curPath) // bingo
		return
	}

	// Golang 中 slice 是传址引用，必须 copy 后使用
	lPath := make([]int, len(curPath))
	copy(lPath, curPath)
	rPath := make([]int, len(curPath))
	copy(rPath, curPath)

	traverse(root.Left, curSum, target, lPath, paths)  // recording...
	traverse(root.Right, curSum, target, rPath, paths) // 一个节点一条路
}
