package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
}

// 哈哈哈评论吐槽语体教
// 静下心来读题 "按照上一个数，报数，得到下一个数"，报数规则与 1 2 3 不同罢了
func countAndSay(n int) string {
	tmp := "1"
	for i := 1; i < n; i++ {
		tmp = next(tmp)
	}
	return tmp
}

func next(s string) (res string) {
	nums := strings.Split(s, "")
	for i := 0; i < len(nums); {
		count := 1
		n := nums[i]
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			count++
			i++
		}
		res += fmt.Sprintf("%d%s", count, n) // 计数连数字
		i++
	}
	return res
}
