package main

import "fmt"

func main() {
	fmt.Println(maxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
}

//
// 暴力遍历
//
func maxArea1(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	maxMulti := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			w := j - i                     // 宽
			h := min(height[i], height[j]) // 高
			maxMulti = max(maxMulti, w*h)  // 比较面积
		}
	}
	return maxMulti
}

//
// 优化1：双指针法
// 两条线段之间的面积受限与最短的线段，线段间距越长，面积越大
// 使用 2 个指针指向首部和尾部，将短指针向长指针方向移动，看能不能找到更长的线，使面积更大
// 依据：向长线方向每次移动 1 格，虽然宽度-1，但是(高度变高)*(宽度-1) >= 高度*宽度
//
func maxArea2(height []int) int {
	maxMulti := 0
	left, right := 0, len(height)-1
	for left < right {
		w := right - left
		h := min(height[left], height[right])
		maxMulti = max(maxMulti, w*h)
		if height[left] <= height[right] {
			left++ // 往右边走找更长的线
		} else {
			right-- // 往左边走
		}
	}
	return maxMulti
}
