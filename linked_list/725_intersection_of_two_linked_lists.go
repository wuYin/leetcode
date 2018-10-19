package main

import "fmt"

func main() {
	for _, node := range splitListToParts(newList([]int{1, 2, 3, 4, 5, 6, 7}), 3) {
		fmt.Println(node)
	}
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	// case: [], 3 => [],[],[]
	// if root == nil {
	// 	return nil
	// }

	n := 0
	cur := root
	for cur != nil {
		n++
		cur = cur.Next
	}
	nodes := make([]*ListNode, k)
	cur = root

	div, remain := n/k, n%k // 商，余数
	fmt.Println(div, remain)
	for i := 0; i < k && cur != nil; i++ {
		nodes[i] = cur
		// 十分巧妙的写法，若数组未达到余数那么长，前边每部分部分都要多接一个节点
		for j := 1; j < div+bool2int(i < remain); j++ {
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}
	return nodes
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
