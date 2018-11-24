package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 4}
	for _, v := range []int{2, 7, 1, 3, 6, 9} {
		tree.BFSInsert(v)
	}
	fmt.Println(levelOrder(tree.root))
}

// 层次遍历（宽度优先遍历）
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var floor []int
		level := len(queue) // level 表示第几层，用于取出当前层的所有值
		for i := 0; i < level; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			floor = append(floor, cur.Val)
		}
		res = append(res, floor)
	}
	return res
}
