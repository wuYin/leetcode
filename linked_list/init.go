package main

import "fmt"

// go run init.go xxx.go
type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	cur := head
	for i := 1; i < n; i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		cur.Next = newNode
		cur = newNode
	}
	return head
}

func (cur *ListNode) String() string {
	counts := 0
	var nums []int
	for cur != nil {
		nums = append(nums, cur.Val)
		counts++
		cur = cur.Next
	}
	return fmt.Sprintf("%d nodes: %v", counts, nums)
}
