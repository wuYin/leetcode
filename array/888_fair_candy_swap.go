package main

import (
	"fmt"
)

func main() {
	fmt.Println(fairCandySwap([]int{1, 2, 5}, []int{2, 4})) // [5,4]
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))    // [2,1]
	fmt.Println(fairCandySwap([]int{32, 38, 8, 1, 9}, []int{63, 92}))
}

// 交换 a,b 后 sumA==sumB，和大的数组相当于捡了芝麻丢了西瓜，有个两倍的关系在里边
func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := sum(A), sum(B)
	diff := abs(sumA-sumB) / 2

	m := make(map[int]int)
	for _, b := range B {
		m[b]++
	}

	for _, a := range A {
		target := a - diff
		if sumA < sumB {
			target = a + diff
		}
		if _, ok := m[target]; ok {
			return []int{a, target}
		}
	}
	return nil
}

func sum(nums []int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
