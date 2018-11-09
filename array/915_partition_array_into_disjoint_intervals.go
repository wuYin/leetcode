package main

import "fmt"

func main() {
	fmt.Println(partitionDisjoint([]int{5, 0, 3, 8, 6})) // 3
}

// 按题的描述
// 用两个 max 变量记录左侧最大值和数组最大值，类似快慢指针的思路
func partitionDisjoint(A []int) int {
	leftMax, curMax, edge := A[0], A[0], 1
	for i, a := range A {
		// fmt.Println("a", a, "leftMax", leftMax, "curMax", curMax)
		if leftMax > a { // 比左侧最大值还小，左小数组向右拉长
			leftMax = curMax // 目前的最大值同步到 leftMax
			edge = i + 1
		}
		if a > curMax {
			curMax = a
		}
	}
	return edge
}
