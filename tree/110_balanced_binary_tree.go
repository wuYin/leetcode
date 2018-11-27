package main

func main() {
	// debug in LeetCode Code editor
}

// 递归从下往上判断各子树的平衡结果
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lDepth := depth(root.Left)
	rDepth := depth(root.Right)
	if abs(lDepth-rDepth) > 1 { // 平衡条件
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 某子树的深度
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := depth(root.Left) + 1
	rDepth := depth(root.Right) + 1 // 加节点本身的深度

	if lDepth > rDepth {
		return lDepth
	}
	return rDepth
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
