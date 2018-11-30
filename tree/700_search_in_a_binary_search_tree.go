package main

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}

	//        4
	//       / \
	//      2   7
	//     / \
	//    1   3
	fmt.Println(searchBST(root, 2)) // 2
}

// what can I say
func searchBST(root *TreeNode, val int) *TreeNode {
	if root != nil {
		switch {
		case root.Val < val:
			return searchBST(root.Right, val)
		case root.Val > val:
			return searchBST(root.Left, val)
		default:
			return root
		}
	}
	return nil
}
