package main

func main() {
	// fmt.Println(generateTrees(3))
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	var roots []*TreeNode

	switch {
	case start > end: // 某一边子树无节点，左斜树或右斜树
		roots = append(roots, nil)

	case start == end: // 某一边子树只有一个节点
		roots = append(roots, &TreeNode{Val: start})

	case start < end:
		for rootV := start; rootV <= end; rootV++ { // 以 rootV 为根节点
			leftTrees := generate(start, rootV-1)
			rightTrees := generate(rootV+1, end)

			// 组合左右子树
			for _, ltree := range leftTrees {
				for _, rtree := range rightTrees {
					root := &TreeNode{ // 以 rootV 为根节点的一棵 BST
						Val:   rootV,
						Left:  ltree,
						Right: rtree,
					}
					roots = append(roots, root)
				}
			}
		}
	}

	return roots
}
