package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}

	//    1
	//    \
	//     2
	//    /
	//   2
	fmt.Println(findMode(root)) // 2
}

// 辣鸡解法是将二叉树随便遍历到计数 map 后，两次遍历求众数即可
// cool solution // 理解本质
func findMode(root *TreeNode) []int {
	var res []int
	var cur, count, max int
	traverseBST(root, &cur, &count, &max, &res)
	return res
}

// 既然是 BST，那就中序遍历取出有序序列，一边遍历一边比较并取众数
func traverseBST(root *TreeNode, cur, count, max *int, res *[]int) {
	if root == nil {
		return
	}
	traverseBST(root.Left, cur, count, max, res)

	if root.Val != *cur {
		*cur = root.Val // 新值
		*count = 1
	} else {
		*count++
	}

	// 新众数纪录
	if *count > *max {
		*max = *count
		*res = nil // 重来
	}

	// 符合众数条件
	if *count == *max {
		*res = append(*res, root.Val)
	}
	traverseBST(root.Right, cur, count, max, res)
}
