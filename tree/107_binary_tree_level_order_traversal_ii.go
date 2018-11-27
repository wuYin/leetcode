package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 3}
	for _, v := range []int{9, 20, 15, 7} {
		tree.BFSInsert(v)
	}
	fmt.Println(levelOrderBottom(tree.root))
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var floor []int
		level := len(queue)

		for i := 0; i < level; i++ {
			cur := queue[0]
			queue = queue[1:]
			floor = append(floor, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		res = append(res, floor)
	}

	// 反转即可
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
