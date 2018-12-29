package main

import "fmt"

func main() {
	fmt.Println(integerBreak(2))  // 1
	fmt.Println(integerBreak(5))  // 6
	fmt.Println(integerBreak(10)) // 36
}

// 和 70 题相似，但需要检查并嵌套两层循环
// 注意 bottom-up 向上存储值时，要取 maxP 和 i 的最大值
func integerBreak(n int) int {
	ps := make([]int, n+1)
	ps[1] = 1 // 边界值

	for i := 2; i <= n; i++ {
		var maxP int
		for j := 1; j <= i/2; j++ {
			maxP = max(maxP, ps[j]*ps[i-j]) // 状态转移
		}
		if i == n {
			return maxP // bingo
		}

		ps[i] = max(maxP, i) // 最优子结构 // 存储中间计算结果
	}

	return -1 // 代码不会执行到这
}
