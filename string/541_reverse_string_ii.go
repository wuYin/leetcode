package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcdefg", 2)) // bacdfeg
}

// 逻辑...
// 画区间图就简洁明了
func reverseStr(s string, k int) string {
	runes := []rune(s)
	n := len(runes)

	reverse := func(rs []rune) []rune {
		s, e := 0, len(rs)-1
		for s <= e {
			rs[s], rs[e] = rs[e], rs[s]
			s++
			e--
		}
		return rs
	}

	res := make([]rune, 0, len(runes))
	switch {
	case n < k:
		res = append(res, reverse(runes[:])...)
	case n < 2*k:
		res = append(res, reverse(runes[:k])...)
	default:
		segs := n / (2 * k)
		start := 0
		for i := 0; i < segs; i ++ {
			start = 2 * i * k
			res = append(res, reverse(runes[start:start+k])...) // 不断反转 [even*k, even*k+k-1] 偶数起始区间的字符
		}
		start += 2 * k
		remain := n - start
		switch {
		case remain < k:
			res = append(res, reverse(runes[start:])...)
		case remain < 2*k:
			res = append(res, reverse(runes[start:start+k])...)
		}
	}
	return string(runes)
}
