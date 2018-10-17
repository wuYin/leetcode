package main

import "fmt"

func main() {
	// fmt.Println(deleteDuplicates(newList([]int{1, 2, 3, 3, 4, 4, 5})))
	fmt.Println(deleteDuplicates(newList([]int{1, 1, 1, 2, 2, 3})))
	// fmt.Println(deleteDuplicates(newList([]int{1, 2, 2})))
}

// 遇到头节点也需要遍历，而且头节点有可能丢失的情况下，一定要求助于 "哑节点"
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil {
		last := pre.Next
		for last.Next != nil && last.Val == last.Next.Val {
			last = last.Next
		}
		// 对比每个序列的第一个数和最后一个数（节点）是否相等
		if pre.Next != last {
			pre.Next = last.Next // 不相等则 pre.Next 指向下一个数字序列的第一个节点
		} else {
			pre = pre.Next // 相等则 ok
		}
	}

	return dummy.Next
}
