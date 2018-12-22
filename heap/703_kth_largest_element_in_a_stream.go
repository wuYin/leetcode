package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 5, 8, 2}
	kthLargest := Constructor(3, arr)
	fmt.Println(kthLargest.Add(3))  // 4
	fmt.Println(kthLargest.Add(5))  // 5
	fmt.Println(kthLargest.Add(10)) // 5
	fmt.Println(kthLargest.Add(9))  // 8
	fmt.Println(kthLargest.Add(4))  // 8
}

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{k: k, nums: nums}
}

func (this *KthLargest) Add(val int) int {
	insert(&this.nums, val) // 低效
	return this.nums[len(this.nums)-this.k]
}

func insert(nums *[]int, v int) {
	ns := *nums
	l, r := 0, len(ns)-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case v < ns[mid]:
			r = mid - 1
		case v >= ns[mid]:
			l = mid + 1
		}
	}
	*nums = append(ns[:l], append([]int{v}, ns[l:]...)...)
}
