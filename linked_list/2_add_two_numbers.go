package main

import "fmt"

// go run init.go 2_add_two_numbers.go
func main() {
	l1 := newList([]int{8, 9, 9}) // 5	// 9 8
	l2 := newList([]int{2})       // 5	// 1
	cur := addTwoNumbers(l1, l2)
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
}


// 遍历两个链表，取出各自的数字再相加（难以解决整数溢出的问题，1560 / 1563 个通过测试用例，取值求解不可行）
// 遍历链表，相互相加记进位。注意特殊情况的处理
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nums []int
	cur1, cur2 := l1, l2
	var carryBit bool // 是否要进位
	for cur1 != nil && cur2 != nil {
		sum := cur1.Val + cur2.Val
		if carryBit {
			sum++
		}
		carryBit = false
		if sum >= 10 {
			carryBit = true
		}
		nums = append(nums, sum%10)
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	nums = append(nums, traverse(cur1, carryBit)...)
	nums = append(nums, traverse(cur2, carryBit)...)
	if cur1 == nil && cur2 == nil && carryBit {
		nums = append(nums, 1)
	}
	return newList(nums)
}

func traverse(cur *ListNode, carryBit bool) (remainNums []int) {
	if cur == nil {
		return
	}
	for cur != nil {
		if carryBit {
			res := cur.Val + 1
			if res >= 10 {
				remainNums = append(remainNums, 0)
				carryBit = true
			} else {
				remainNums = append(remainNums, res)
				carryBit = false
			}
			cur = cur.Next
			continue
		}
		remainNums = append(remainNums, cur.Val)
		cur = cur.Next
	}
	if carryBit {
		remainNums = append(remainNums, 1)
	}
	return
}

