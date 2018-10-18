package main

import "fmt"

func main() {
	fmt.Println(numComponents(newList([]int{0, 1, 2, 3}), []int{0, 1, 3}))
	fmt.Println(numComponents(newList([]int{0, 1, 2, 3, 4}), []int{0, 1, 3, 4}))
}

// 想起了两数之和，遍历链表时判断值是否在 G 中，使用散列表 map 就很好
// 注意理清楚逻辑就行
func numComponents(head *ListNode, G []int) int {
	if head == nil || len(G) == 0 {
		return 0
	}
	nums := make(map[int]int)
	for _, g := range G {
		nums[g]++
	}
	counts := 0
	front, rear := head, head
	for rear != nil {
		if _, ok := nums[front.Val]; !ok {
			front = front.Next
			rear = rear.Next
			continue
		}
		// 至少找到一个存在的值
		counts++
		for rear != nil {
			if _, ok := nums[rear.Val]; !ok {
				break
			}
			rear = rear.Next // 跳过所有连续、在 G 中有的值
		}
		front = rear
	}
	return counts
}
