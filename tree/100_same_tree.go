package main

func main() {
	// just bingo
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return same(p, q)
}

func same(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil { // 先 && 再 || 相当于 1+2=3 种判断
		return false
	}
	return node1.Val == node2.Val && same(node1.Left, node2.Left) && same(node1.Right, node2.Right)
}
