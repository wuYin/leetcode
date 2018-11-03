package main

import "fmt"

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))       // 2
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})) // 6
}


// 辣鸡解法
func findShortestSubArray(nums []int) int {
	degrees := make(map[int]int)
	for _, num := range nums {
		degrees[num]++
	}
	maxDegree := 1
	var targets []int
	for num, degree := range degrees {
		if degree >= maxDegree {
			maxDegree = degree
			targets = append(targets, num)
		}
	}

	// fmt.Println("maxDegree", maxDegree)
	minLen := len(nums)
	for _, target := range targets {
		for i, num := range nums {
			if num == target {
				subLen := 1
				numCount := 1
				for j := i + 1; j < len(nums); j++ {
					subLen++
					if nums[j] == target {
						numCount++
					}
					if numCount >= maxDegree {
						break
					}
				}
				// fmt.Println(num, numCount, maxDegree, subLen, minLen)
				if numCount >= maxDegree && subLen < minLen {
					minLen = subLen
				}
			}
		}
	}
	return minLen
}
