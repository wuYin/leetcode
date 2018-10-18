package main

import "fmt"

func main() {
	fmt.Println(oddEvenList(newList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))
}

// 小样本发现不了规律，就把数据量设置大一点，走一遍题目要求，就看出规律来了
// 1 	2 	3 		4 5 6 7 8 9
// pre 	cur target	// 不断交换 cur 和 target 的值即可，注意二者的 gap 间距从 1 开始递增
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	gap := 1
	pre, cur, target := head, head.Next, head.Next.Next
	for {
		swap3Nodes(pre, cur, target)
		gap++
		for i := 0; target != nil && i <= gap; i++ {
			target = target.Next
		}
		pre = pre.Next
		if target == nil {
			break
		}
	}
	return head
}

// 交换 pre->cur->...->target->... 三个节点的值
func swap3Nodes(pre, cur, target *ListNode) {
	pre.Next = target
	targetNext := target.Next
	target.Next = cur
	curNext := cur
	for curNext.Next != target {
		curNext = curNext.Next
	}
	curNext.Next = targetNext
}
