package main

import "fmt"

func main() {
	s := MinStack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Top())    // 2
	fmt.Println(s.GetMin()) // 1
	s.Pop()
	fmt.Println(s.GetMin()) // 1
	fmt.Println(s.Top())    // 1
}

// 设计具有 GetMin 操作的最小栈
type MinStack struct {
	min   int
	stack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.min = x
	}
	if x < this.min {
		this.min = x
	}
	this.stack = append(this.stack, x)
}

// pop 操作后需要解决 min 的操作
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.stack) > 0 {
		this.min = this.stack[0]
	}
	for _, elem := range this.stack {
		if this.min > elem {
			this.min = elem
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
