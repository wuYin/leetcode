package main

import "fmt"

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
}

// 可以直接使用 []int 来实现 queue
// 比较取巧的方法，用两个栈实现一个队列: x->s1,s2->
type MyQueue struct {
	s1, s2 Stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() { // 所有元素向前挪动
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() { // 所有元素向前挪动
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.s2.IsEmpty() && this.s1.IsEmpty()
}
