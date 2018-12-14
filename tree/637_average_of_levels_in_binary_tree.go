package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20},
	}

	fmt.Println(averageOfLevels(root)) // [3 14.5]
}

// 层序遍历即可
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var avgs []float64
	q := []*TreeNode{root}
	for len(q) > 0 {
		count := len(q)
		sum := 0
		for i := 0; i < count; i++ {
			cur := q[0]
			q = q[1:]
			sum += cur.Val
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		avgs = append(avgs, float64(sum)/float64(count))
	}
	return avgs
}
