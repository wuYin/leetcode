package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(newList([]int{1, 2})))
	fmt.Println(isPalindrome(newList([]int{1, 2, 2, 1})))
	fmt.Println(isPalindrome(newList([]int{1, 2, 3, 2, 1})))
}

// 将前半段链表反转到新链表，再逐一对比新链表与后半段链表是否一致
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true // 注意 []、[1] 是回文结构
	}

	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	mid := n / 2 // 找到中间节点

	cur = head
	newHead := &ListNode{Val: head.Val} // 遍历新建前半段链表
	i := 0
	for i < mid-1 {
		cur = cur.Next
		newHead = prepend(newHead, cur.Val)
		i++
	}

	cur = cur.Next // 选取要开始遍历的后半段链表
	if n%2 == 1 {
		cur = cur.Next
	}

	for newHead != nil && cur != nil {
		if newHead.Val != cur.Val {
			return false
		}
		cur = cur.Next
		newHead = newHead.Next
	}

	return true // 对称
}

// 新增 head
func prepend(head *ListNode, v int) (newHead *ListNode) {
	node := &ListNode{Val: v}
	if head == nil {
		return node
	}
	node.Next = head
	return node
}
