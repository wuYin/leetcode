package main

import (
	"fmt"
	"sort"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 9}

	// 	        1
	//         / \
	//        3   2
	//       / \   \
	//      5   3   9
	fmt.Println(largestValues(root)) // [1 3 9]
}

func largestValues(root *TreeNode) []int {
	var levels [][]int
	var traverse func(root *TreeNode, curLevel int)

	// 和 102 一样的层序遍历
	// 递归和 queue 都 ok
	traverse = func(root *TreeNode, curLevel int) {
		if root == nil {
			return
		}

		// 到了新一层
		if curLevel >= len(levels) {
			levels = append(levels, []int{})
		}
		// 注意此处 levels 直接使用的外部变量，传 *[][]int 参数无法 append
		levels[curLevel] = append(levels[curLevel], root.Val) // bingo

		traverse(root.Left, curLevel+1)
		traverse(root.Right, curLevel+1)
	}

	traverse(root, 0)

	var res []int
	for _, level := range levels {
		sort.Ints(level)
		res = append(res, level[len(level)-1])
	}
	return res
}
