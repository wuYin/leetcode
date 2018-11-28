package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}

	//    1            <---
	//  /   \
	// 2     3         <---
	//  \     \
	//   5     4       <---
	fmt.Println(rightSideView(root)) // [1 3 4]
}

// 从右侧看，其实是要求每层的最右侧节点
// 层序遍历每层取最后一个元素即可
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var floors [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var floor []int      // 本层的节点值
		counts := len(queue) // 本层的节点数
		for i := 0; i < counts; i++ {
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
		floors = append(floors, floor)
	}

	var nums []int
	for _, floor := range floors {
		nums = append(nums, floor[len(floor)-1]) // bingo
	}
	return nums
}
