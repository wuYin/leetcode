package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
}

type Trie struct {
	IsWord bool
	Next   map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		IsWord: false,
		Next:   make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, r := range word {
		if next, ok := cur.Next[r]; ok {
			cur = next
		} else {
			if cur.Next == nil {
				cur.Next = make(map[rune]*Trie)
			}
			newNode := &Trie{IsWord: false, Next: nil}
			cur.Next[r] = newNode
			cur = newNode
		}
	}
	cur.IsWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word) && cur != nil; i++ {
		if next, ok := cur.Next[rune(word[i])]; !ok {
			return false
		} else {
			cur = next
		}
	}
	return cur != nil && cur.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix) && cur != nil; i++ {
		if next, ok := cur.Next[rune(prefix[i])]; !ok {
			return false
		} else {
			cur = next
		}
	}
	return cur != nil
}
