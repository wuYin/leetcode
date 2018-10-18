package main

import "fmt"

func main() {
	fmt.Println(removeElements(newList([]int{1, 1}), 1))
	fmt.Println(removeElements(newList([]int{1, 2, 6, 3, 4, 5, 6}), 6))
}

// 头节点可能会被删掉，要考虑借助哑节点
// 画图理解就 ok
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		pre = pre.Next
		cur = cur.Next
	}
	return dummy.Next
}
