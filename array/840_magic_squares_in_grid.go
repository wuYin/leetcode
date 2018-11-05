package main

import "fmt"

func main() {
	fmt.Println(numMagicSquaresInside([][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}))

}

// 1~9 去填充 9 个格子，横、竖、对角线三个方向的和要一致，中间值必须为 5，其余方向的两个数和必须为 10，总和必须为 15
// 代码也是绕得慌...
func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 {
		return 0
	}

	r, c := len(grid)-1, len(grid[0])-1
	counts := 0
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if grid[i][j] != 5 {
				continue
			}

			ok := true
			m := make([]bool, 9)
			m[4] = true
			offsets := [][]int{{-1, -1}, {-1, 0}, {0, -1}, {-1, 1}} // 八个方向

			for _, offset := range offsets {
				v1 := grid[i+offset[0]][j+offset[1]]
				v2 := grid[i-offset[0]][j-offset[1]]
				if v1 < 1 || v1 > 9 || v2 < 1 || v2 > 9 || v1+v2 != 10 || m[v1-1] || m[v2-1] {
					ok = false
					break
				}
				m[v1-1], m[v2-1] = true, true
			}
			if ok {
				counts++
			}
		}
	}
	return counts
}
