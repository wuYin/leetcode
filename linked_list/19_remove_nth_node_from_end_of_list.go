package main

import "fmt"

func main() {
	fmt.Println(removeNthFromEnd(newList([]int{1, 2}), 1))
	// fmt.Println(removeNthFromEnd(newList([]int{1, 2}), 2))
}

// 和倒数相关的问题考虑天生有间距的"双指针"，一次遍历解决
// 借助哑节点是关键，用于防止删除倒数第 len(nums) 个节点会导致链表丢失
// 若要兼顾正常节点的移动，又要处理头结点的特殊情况，请考虑"哑节点"
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	front, rear := dummyNode, dummyNode
	for counts := 0; counts <= n; counts++ {
		rear = rear.Next
	}
	for rear != nil {
		front = front.Next
		rear = rear.Next
	}
	front.Next = front.Next.Next // 删除节点
	return dummyNode.Next
}
