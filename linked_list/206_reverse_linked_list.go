package main

import "fmt"

func main() {
	fmt.Println(reverseList(newList([]int{1, 2, 3, 4, 5})))
}

// 有点栈的味道哦
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Val: head.Val}
	cur := head.Next
	for cur != nil {
		newHead = prepend(newHead, cur.Val)
		cur = cur.Next
	}
	return newHead
}

func prepend(head *ListNode, v int) *ListNode {
	node := &ListNode{Val: v}
	if head == nil {
		return node
	}
	node.Next = head
	return node
}
