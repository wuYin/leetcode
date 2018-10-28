package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{1, 2})) // 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))    // 2.5
}

//
// 合并两个有序数组再排序，分数组长度情况取中位数即可
//
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	// nums := append(nums1, nums2...)
	// sort.Ints(nums) 			// 快排的复杂度为 O(NlogN)
	nums := mergeArr(nums1, nums2) // 归并排序的 merge 操作复杂度为 O(N)
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		return float64(nums[mid]+nums[mid-1]) / 2
	}
	return float64(nums[mid])
}

//
// 优化1
// 合并有序数组后仍有序，自然而然应该想到归并排序的 merge 操作
//
func mergeArr(arr1, arr2 []int) []int {
	n1, n2 := len(arr1), len(arr2)
	res := make([]int, 0, n1+n2)
	j, k := 0, 0
	for i := 0; i < n1+n2; i++ {
		// 若两个数组长度不一，需避免出现溢出
		if j >= n1 {
			res = append(res, arr2[k:]...)
			break
		}
		if k >= n2 {
			res = append(res, arr1[j:]...)
			break
		}
		// 正常
		if arr1[j] < arr2[k] {
			res = append(res, arr1[j])
			j++
		} else {
			res = append(res, arr2[k])
			k++
		}
	}
	return res
}

//
// 优化2
// 将问题的本质抽象出来：在两个有序数组中寻找第 K 大的数
// O(logN) 的高效查找想到二分查找，两个有序数组的高效查找也是这个思路
//
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	c := len(nums1) + len(nums2)
	if c%2 == 1 {
		return float64(findKth(nums1, nums2, c/2))
	}
	return (float64(findKth(nums1, nums2, c/2-1)) + float64(findKth(nums1, nums2, c/2))) / 2
}
func findKth(nums1, nums2 []int, k int) int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	if len(nums1) == 0 {
		return nums2[k]
	}
	if k == len(nums1)+len(nums2)-1 {
		return max(nums1[len(nums1)-1], nums2[len(nums2)-1])
	}

	i := min(len(nums1)-1, k/2)
	j := min(len(nums2)-1, k-i)
	if nums1[i] < nums2[j] {
		return findKth(nums1[i:], nums2[:j], j)
	}
	return findKth(nums1[:i], nums2[j:], i)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
