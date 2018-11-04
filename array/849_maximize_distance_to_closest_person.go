package main

import "fmt"

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))                               // 2
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))                                        // 3
	fmt.Println(maxDistToClosest([]int{0, 1}))                                              // 1
	fmt.Println(maxDistToClosest([]int{0, 1, 1, 1, 0, 0, 1, 0, 0}))                         // 2
	fmt.Println(maxDistToClosest([]int{0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0})) // 3
}

// 找到最长的 0 序列，按序列边界是否为 1 来决定取中点还是边界
// 双指针
// 辣鸡解法
func maxDistToClosest(seats []int) int {
	n := len(seats)
	if n <= 0 {
		return n
	}

	// 0 开头或 0 结尾的两种特殊 case
	maxStart, maxEnd := 0, 0
	if seats[0] == 0 {
		for i := 0; i < n && seats[i] == 0; i++ {
			maxStart++
		}
	}
	if seats[n-1] == 0 {
		for i := n - 1; i >= 0 && seats[i] == 0; i-- {
			maxEnd ++
		}
	}

	// 快慢指针找出位于中间的最长 0 序列
	slow := 0
	zeroEdges := make([]int, 2)
	for slow < n {
		fast := slow
		for fast < n && seats[fast] == 0 {
			fast++
		}
		if fast-1 > slow {
			if fast-1-slow > zeroEdges[1]-zeroEdges[0] {
				zeroEdges[0], zeroEdges[1] = slow, fast-1
			}
		}
		slow++
	}

	return max(maxStart, maxEnd, (zeroEdges[1]-zeroEdges[0])/2+1)
}

func max(x, y, z int) int {
	if x > y {
		if x > z {
			return x
		}
		return z
	} else {
		if y > z {
			return y
		}
		return z
	}
}
