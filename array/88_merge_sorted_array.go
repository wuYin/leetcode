package main

import "fmt"

func main() {
	nums1 := []int{2}
	nums2 := []int{}
	merge(nums1, 1, nums2, 0)
	fmt.Println(nums1)
}

// 两个指针指向两个数组，分别向后遍历
// 注意解决未遍历完的数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	cpNums := make([]int, 0, m+n)
	copy(cpNums, nums1)

	p1, p2 := 0, 0
	nums := nums1[:m]
	for p1 < m && p2 < n {
		if nums[p1] < nums2[p2] {
			cpNums = append(cpNums, nums[p1])
			p1++
		} else {
			cpNums = append(cpNums, nums2[p2])
			p2++
		}
	}

	if p1 < m {
		cpNums = append(cpNums, nums[p1:]...)
	}
	if p2 < n {
		cpNums = append(cpNums, nums2[p2:]...)
	}

	for i, num := range cpNums {
		nums1[i] = num // 函数内部修改 slice 只能之指定索引修改，与指定传指针原理相同
	}
}
