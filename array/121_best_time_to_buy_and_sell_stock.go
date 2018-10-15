package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1,2}))
}

// 简单的规律
func maxProfit(prices []int) int {
	n := len(prices)
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
