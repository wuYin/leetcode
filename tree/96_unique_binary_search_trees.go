package main

import "fmt"

func main() {
	fmt.Println(numTrees(3)) // 5
}

// 有点类似斐波那契数列的味道
// 1 -> 1
// 2 -> 2
// 3 -> 1*2 + 1*1 + 2*1 = 5
// 4 -> 1*5 + 2*2 + 5*1 = 14
// ...
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1] // -1 是剔除本体 // 以 i 为根节点的 BST 数 = 两侧子 BST 数乘积
		}
	}
	return dp[n]
}
