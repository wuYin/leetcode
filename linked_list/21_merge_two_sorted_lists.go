package main

import "fmt"

func main() {
	fmt.Println(mergeTwoLists(newList([]int{1}), newList([]int{})))
}

// 借助哑节点，万事大吉
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	dummy := cur
	cur1, cur2 := l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	traverse := func(l, remainL *ListNode) {
		for remainL != nil {
			l.Next = remainL
			l = l.Next
			remainL = remainL.Next
		}
	}
	traverse(cur, cur1)
	traverse(cur, cur2)
	return dummy.Next
}
