package main

import (
	"container/list"
	"fmt"
)

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)           // 2 作废
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)           // 1 作废
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4
}

type CacheItem struct {
	k, v int
}

type LRUCache struct {
	limit int
	items *list.List
	cache map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		limit: capacity,
		items: list.New(),
		cache: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Put(key, value int) {
	// 命中
	if elem, ok := this.cache[key]; ok {
		this.items.MoveToFront(elem)
		elem.Value.(*CacheItem).v = value // 更新缓存值
		return
	}

	// 未命中
	newItem := &CacheItem{k: key, v: value}
	newElem := this.items.PushFront(newItem)
	this.cache[key] = newElem

	// 超过缓存容量
	if this.items.Len() > this.limit {
		oldest := this.items.Back()
		delete(this.cache, oldest.Value.(*CacheItem).k) // 移除最少使用的 item
		this.items.Remove(oldest)
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok { // 命中
		this.items.MoveToFront(elem)
		return elem.Value.(*CacheItem).v
	}

	return -1
}
