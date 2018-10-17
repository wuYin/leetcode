package main

import "fmt"

func main() {
	fmt.Println(rotateRight(newList([]int{1, 2, 3, 4, 5}), 2))
	fmt.Println(rotateRight(newList([]int{0, 1, 2}), 4))
}

// 和 189 类似，注意先取余
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := count(head)
	step := k % n
	for step > 0 {
		pre, tail := tailNodes(head) // 获取到倒数二个节点：pre->tail->nil
		pre.Next = nil               // 删除 tail
		tail.Next = head             // 将 tail 作为新的 head
		head = tail
		step--
	}
	return head
}

func tailNodes(head *ListNode) (*ListNode, *ListNode) {
	pre, tail := head, head.Next
	for tail.Next != nil {
		pre = tail
		tail = tail.Next
	}
	return pre, tail
}

func count(head *ListNode) int {
	counts := 0
	cur := head
	for cur != nil {
		counts++
		cur = cur.Next
	}
	return counts
}
