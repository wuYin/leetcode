package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B'}, 2))
}

// 思路：从高频字母开始
// A _ _ A _ _ A _ _
// A B _ A B _ A _ _
// A B _ A B _ A		// bingo
func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]int)
	for _, t := range tasks {
		m[t]++
	}

	counts := make(map[int][]byte)
	for b, count := range m {
		counts[count] = append(counts[count], b) // 统计各字母的出现频率
	}

	var freqs []int
	for c := range counts {
		freqs = append(freqs, c)
	}
	sort.Ints(freqs)

	maxFreq := freqs[len(freqs)-1]
	c := maxFreq * (n + 1)
	holders := make([]byte, c) // 占位 // 注意遍历顺序是不定的
	for i := range holders {
		if i%(n+1) == 0 {
			holders[i] = counts[maxFreq][0] // 初始化占位数组
		}
	}
	counts[maxFreq] = counts[maxFreq][1:]

	// 从后向前遍历字符，填充
	var cs []int
	for count := range counts {
		cs = append(cs, count)
	}
	sort.Ints(cs)

	// debug(holders)

	for i := len(cs) - 1; i >= 0; i-- {
		c = cs[i]
		bs := counts[c]
		for _, b := range bs {
			for start, h := range holders {
				if h == 0 { // 向后占位
					for i := start; c >= 0 && i < len(holders); c-- {
						holders[i] = b
						i += n + 1
					}
					break
				}
			}
		}
	}

	holders = trimRight(holders)
	debug(holders) // A B 0 A B 0 A B

	return len(holders) // 8
}

func trimRight(bs []byte) []byte {
	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == byte(0) {
			bs = bs[:i]
		} else {
			break
		}
	}
	return bs
}

func debug(bs []byte) {
	s := string(bs)
	s = strings.Replace(s, "\x00", "0", -1)
	fmt.Println(strings.Join(strings.Split(s, ""), " "))
}
