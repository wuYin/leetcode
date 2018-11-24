package main

func main() {
	// just bingo
}

func isSymmetric(root *TreeNode) bool {
	return mirror(root, root)
}

// 镜像树：递归交换左右子树
func mirror(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	// 左子树与右子树值一致
	return node1.Val == node2.Val && mirror(node1.Left, node2.Right) && mirror(node1.Right, node2.Left)
}
