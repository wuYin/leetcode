package main

import "fmt"

func main() {
	fmt.Println(swapPairs(newList([]int{1})))
	fmt.Println(swapPairs(newList([]int{1, 2})))
	fmt.Println(swapPairs(newList([]int{1, 2, 3, 4})))
}

// 处理好 2 种异常 case
// 将两两节点视为内部短链表，交换两个节点后返回新的头节点（原第二节点），接到尾巴上即可
// 一般 case：pre->cur->next 经 swap() 调换后 next->pre，手动接上 pre->next->cur
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 异常case: [] 或 [1]
		return head
	}

	newHead := swap(head, head.Next)
	pre := head
	if pre.Next == nil || pre.Next.Next == nil { // 异常case: [1, 2]
		return newHead
	}

	cur := pre.Next
	next := cur.Next
	for cur != nil && next != nil { // 一般case: 两两结对向后遍历
		pre.Next = swap(cur, next) // 内部交换
		if cur.Next == nil {
			break
		}
		pre = cur
		cur = cur.Next
		next = cur.Next
	}
	return newHead
}

func swap(cur, next *ListNode) *ListNode {
	cur.Next = next.Next
	next.Next = cur
	return next
}
