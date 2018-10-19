package main

import "fmt"

func main() {
	head := newList([]int{1, 2, 3})
	fmt.Println(hasCycle(head)) // false
	head.Next.Next.Next = head.Next
	fmt.Println(hasCycle(head)) // true
}

// map 是保证集合唯一性很好的选择
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	addresses := make(map[*ListNode]int) // 指针类型（内存地址）是可比较的哟
	cur := head
	for cur != nil {
		if _, ok := addresses[cur]; ok {
			return true
		}
		addresses[cur] = 0
		cur = cur.Next
	}
	return false
}
