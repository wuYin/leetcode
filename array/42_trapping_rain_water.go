package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
}

func trap(height []int) int {
	area := 0
	l, r := 0, len(height)-1
	lTop, rTop := 0, 0

	for l <= r {
		lTop = max(height[l], lTop)
		rTop = max(height[r], rTop)
		if lTop < rTop {
			area += lTop - height[l] // 右侧更高些，往右侧走。一边向右移，一边计算高度差来累计面积（最高点处高度差为 0）
			l++
		} else {
			area += rTop - height[r] // 向左侧移，用目前右侧最高高度来计算高度差，累计面积
			r--
		}
	}
	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
