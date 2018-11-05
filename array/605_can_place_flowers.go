package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0}, 1)) // false
}

// 种花的本质是标识 0，即是一边遍历一边修改为 1
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if length == 1 && flowerbed[0] == 0 {
		return true
	}

	// first blood
	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1 // 种花
		n--
	}
	for i := 1; i < length; i++ {
		// 前后都是 0 就可以种花，注意最后一位的 case
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && (i == length-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
