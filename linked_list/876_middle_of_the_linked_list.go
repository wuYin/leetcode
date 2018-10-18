package main

import "fmt"

func main() {
	fmt.Println(middleNode(newList([]int{1, 2, 3, 4, 5})))
	fmt.Println(middleNode(newList([]int{1, 2, 3, 4, 5, 6})))
}

// 考虑遍历的两种普遍问题
// 间距问题：使用2个间距指针 front, rear
// 截取1/N问题：使用2个倍速指针 slow, fast
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 先统计节点个数，再将 head 移动到中间节点即可
func middleNode1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	mid := n / 2
	for i := 0; i < mid; i++ {
		head = head.Next
	}
	return head
}
