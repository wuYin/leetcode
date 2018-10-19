package main

import "fmt"

func main() {
	fmt.Println(reverseBetween(newList([]int{1, 2, 3, 4, 5}), 2, 4))
}

// 更具普遍意义的反转链表，记得借助哑节点以防 head 丢失
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	i := 1
	pre, cur := dummy, head
	for i < m {
		pre = pre.Next
		cur = cur.Next
		i++
	}

	// 开始反转
	newStart := &ListNode{Val: cur.Val}
	for i < n {
		cur = cur.Next
		newStart = prepend(newStart, cur.Val)
		i++
	}
	cur = cur.Next
	pre.Next = newStart

	// 反转完毕，把 [n...] 之后的接上
	for i := 0; i < n-m; i++ {
		newStart = newStart.Next
	}
	newStart.Next = cur

	return dummy.Next
}

func prepend(head *ListNode, v int) *ListNode {
	node := &ListNode{Val: v}
	if head == nil {
		return node
	}
	node.Next = head
	return node
}
