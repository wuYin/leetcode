package main

import "fmt"

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Left.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 6}
	//        1
	//       / \
	//      2   3
	//     /   / \
	//    4   5   6
	//       /
	//      7
	fmt.Println(findBottomLeftValue(root)) // 7
}

// 最左下角的元素，即最底层第一个元素
func findBottomLeftValue(root *TreeNode) int {
	var levels []int
	traverse(root, 0, &levels)
	return levels[len(levels)-1]
}

// 和 199 层序遍历并存储节点值类似
// 需要用到辅助空间存储与长度判断是否到达下一层
func traverse(root *TreeNode, curLevel int, levels *[]int) {
	if root == nil {
		return
	}
	curLevel++
	if curLevel > len(*levels) { // 新一层
		*levels = append(*levels, root.Val)
	}
	traverse(root.Left, curLevel, levels) // 先处理最左节点
	traverse(root.Right, curLevel, levels)
}
