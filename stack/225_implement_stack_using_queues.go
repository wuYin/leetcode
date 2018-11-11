package main

import "fmt"

func main() {
	s := MyStack{}
	fmt.Println(s.Empty()) // true
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop()) // 3
}

// 两个队列实现一个栈
// [...enque...] <-
// [...deque...] ->
type MyStack struct {
	enque, deque []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.enque = append(this.enque, x)
}

func (this *MyStack) Pop() int {
	// 把 enque 队列元素按顺序 append 到 deque 队列中
	n := len(this.enque)
	for i := 0; i < n-1; i++ { // 留最后一个元素
		this.deque = append(this.deque, this.enque[0])
		this.enque = this.enque[1:]
	}
	peek := this.enque[0]
	this.enque = this.deque
	this.deque = nil
	return peek
}

func (this *MyStack) Top() int {
	peek := this.Pop()
	this.enque = append(this.enque, peek)
	return peek
}

func (this *MyStack) Empty() bool {
	return len(this.enque) == 0
}
