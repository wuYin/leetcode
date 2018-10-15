package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{}, 10))
}

// 遍历就完事儿了
func twoSum(numbers []int, target int) []int {
	head, tail := 0, len(numbers)-1
	for head <= tail {
		switch {
		case numbers[head]+numbers[tail] < target:
			head++
		case numbers[head]+numbers[tail] > target:
			tail--
		default:
			return []int{head + 1, tail + 1}
		}
	}
	return nil
}
