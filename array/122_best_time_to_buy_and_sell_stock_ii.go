package main

import "fmt"

func main() {
	fmt.Println(maxProfit1([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit1(prices []int) int {
	i, max := 0, 0
	n := len(prices)
	for i < n-1 {
		buy := i
		for i < n-1 && prices[i] < prices[i+1] {
			i++
		}
		max += prices[i] - prices[buy] // 在最高点抛售
		i++
	}
	return max
}
