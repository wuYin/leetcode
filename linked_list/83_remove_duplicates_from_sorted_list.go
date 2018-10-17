package main

import "fmt"

func main() {
	fmt.Println(deleteDuplicates(newList([]int{1, 1, 2})))
	fmt.Println(deleteDuplicates(newList([]int{1, 1, 2, 3, 3})))
}

// 遍历即可
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}
