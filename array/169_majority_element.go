package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{6, 5, 5}))
	// fmt.Println(majorityElement([]int{3, 2, 3}))
	// fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	major := len(nums) / 2
	for num, times := range m { // map 的 range 顺序不定  ok ?
		if times > major {
			return num
		}
	}
	return 0
}
