package main

import "fmt"

func main() {
	c := Constructor(2)
	c.Put(1, 1)           // 1: 1
	c.Put(2, 2)           // 1: 2->1
	fmt.Println(c.Get(1)) // 1: 1->2
	c.Put(3, 3)           // 2 失效
	fmt.Println(c.Get(2)) // -1
	fmt.Println(c.Get(3)) // 1: 3->1
	c.Put(4, 4)           // 1 失效 // 1: 4->3
	fmt.Println(c.Get(1)) // -1
	fmt.Println(c.Get(3)) // 1: 3->4
	fmt.Println(c.Get(4)) // 1: 4->3
}

type LFUCache struct {
	capacity int
	minFreq  int
	items    map[int]*Node
	freqs    map[int]*List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		minFreq:  0,
		items:    make(map[int]*Node),
		freqs:    make(map[int]*List),
	}
}

func (this *LFUCache) Get(k int) int {
	// 未命中
	node, ok := this.items[k]
	if !ok {
		return -1
	}

	// 命中
	this.freqs[node.freq].Remove(node) // 挪动到下一频率梯队
	node.freq++
	if _, ok := this.freqs[node.freq]; !ok {
		this.freqs[node.freq] = NewList()
	}
	newNode := this.freqs[node.freq].Prepend(node)
	this.items[k] = newNode
	// 注意判断最小梯队
	if this.freqs[this.minFreq].Size() == 0 {
		this.minFreq++
	}
	return node.val
}

func (this *LFUCache) Put(key, value int) {
	if this.capacity == 0 {
		return
	}
	// 命中
	if val := this.Get(key); val != -1 {
		this.items[key].val = value
		return
	}

	// 缓存已满
	if this.capacity == len(this.items) {
		oldest := this.freqs[this.minFreq].Tail()
		this.freqs[this.minFreq].Remove(oldest)
		delete(this.items, oldest.key)
	}

	node := &Node{key: key, val: value, freq: 1}
	this.items[key] = node
	if _, ok := this.freqs[1]; !ok {
		this.freqs[1] = NewList()
	}
	this.freqs[1].Prepend(node)
	this.minFreq = 1
}

type Node struct {
	key        int
	val        int
	freq       int
	prev, next *Node
}

type List struct {
	head, tail *Node
	size       int
}

func NewList() *List {
	return new(List)
}

func (l *List) Prepend(node *Node) *Node {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = nil
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.size++
	return l.head
}

func (l *List) Remove(node *Node) *Node {
	if node == nil {
		return nil
	}

	prev, next := node.prev, node.next
	if prev == nil {
		l.head = next
	} else {
		prev.next = next
	}

	if next == nil {
		l.tail = prev
	} else {
		next.prev = prev
	}

	node.prev, node.next = nil, nil
	l.size--
	return node
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Tail() *Node {
	return l.tail
}
