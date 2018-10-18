package main

import "fmt"

func main() {
	fmt.Println(deleteNode(newList([]int{4, 5, 1, 9}), 5))
	fmt.Println(deleteNode(newList([]int{4, 5, 1, 9}), 9))
	fmt.Println(deleteNode(newList([]int{1, 2}), 1))
}

// 注意 break 条件即可
func deleteNode(head *ListNode, v int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	pre, cur := head, head.Next
	for {
		if cur.Next == nil || cur.Val == v {
			break
		}
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = cur.Next
	return head
}
