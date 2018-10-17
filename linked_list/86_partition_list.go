package main

import "fmt"

func main() {
	fmt.Println(partition(newList([]int{1, 4, 3, 2, 5, 2}), 3))
}

// 哑节点是利器啊
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy1, dummy2 := new(ListNode), new(ListNode)
	cur1, cur2 := dummy1, dummy2
	cur := head
	for cur != nil { // 遍历构建
		if cur.Val < x {
			cur1.Next = cur
			cur1 = cur
		} else {
			cur2.Next = cur
			cur2 = cur
		}
		cur = cur.Next
	}
	cur1.Next = nil // 截断
	cur2.Next = nil
	cur1.Next = dummy2.Next // 连接
	return dummy1.Next
}
