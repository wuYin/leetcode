package main

import "fmt"

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2)) // 4
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2)) // 3
}

// 对我这种不玩LOL的人来说本题的超长描述就像这句注释一样对题解没啥用...
// 题意：求解重叠区间
func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n <= 0 {
		return 0
	}

	t := 0
	for i := 1; i < n; i++ {
		diff := timeSeries[i] - timeSeries[i-1]
		if diff > duration { // 未重叠
			t += duration
		} else {
			t += diff // 重叠部分不予处理
		}
	}
	return t + duration // 最后一个数
}

// 辣鸡的第一反应解法
func findPoisonedDuration2(timeSeries []int, duration int) int {
	m := make(map[int]int)
	for _, t := range timeSeries {
		for i := 0; i < duration; i++ {
			m[t+i] = 0
		}
	}
	return len(m)
}
