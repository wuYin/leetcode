package main

import "fmt"

func main() {
	fmt.Println(badLongestPalindrome("babad"))  // bab
	fmt.Println(goodLongestPalindrome("babad")) // bab
	fmt.Println(bestLongestPalindrome("babad")) // bab
}

//
// Manacher 马拉车算法
// O(N) 检测最长回文子串
//
func bestLongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	t := "#"
	for _, r := range s {
		t += string(r) + "#" // 统一将 偶数串/奇数串 -> 奇数串
	}

	radius := make([]int, len(t)) // 每个位置的回文半径
	maxCenter := 0                // 到目前为止最长回文串的中心位置
	maxR := 0                     // 到目前位置最长回文串的右边界索引
	stopCenter := 0
	maxRadius := 0

	for i := 0; i < len(t); i++ {
		mirror := 2*maxCenter - i
		if i < maxR { // 没有超过最大右边界
			radius[i] = min(radius[mirror], maxR-i) // 取 [i,mx] 与 i'半径 的最小值
		}

		// 以 i 为中心不断扩充半径，向左右两侧探测回文
		// 不要以为这里复杂度为 O(N)
		// 若上边 radius[i] = radius[i']，则不会再进入循环扩张半径
		for i+1+radius[i] < len(t) && i-1-radius[i] >= 0 && t[i-1-radius[i]] == t[i+1+radius[i]] {
			radius[i]++
		}

		// 超出了边缘
		if i+radius[i] > maxR {
			maxR = i + radius[i]
			maxCenter = i
		}

		if radius[i] > maxRadius {
			maxRadius = radius[i]
			stopCenter = i
		}
	}
	l := stopCenter/2 - maxRadius/2 // t->s，索引取中
	r := l + maxRadius - 1          // 减掉自身
	return string(s[l : r+1])       // bingo
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//
// Golang 的题解中不要使用全局作用域变量，OJ 每个 case 不会予以清零
// O(N^2) 复杂度
//
func goodLongestPalindrome(s string) string {
	runes := []rune(s)
	var maxLen int
	var maxSubStr string
	for i := range runes {
		length, str := spread(runes, i, i, maxLen) // aba
		if length > maxLen {
			maxLen = length
			maxSubStr = str
		}
		length, str = spread(runes, i, i+1, maxLen) // abba
		if length > maxLen {
			maxLen = length
			maxSubStr = str
		}
	}
	return maxSubStr
}

// 以中心点向两侧扩散，获取当前字符为中心的最大回文串
func spread(runes []rune, l, r int, curMaxLen int) (length int, subStr string) {
	length = curMaxLen
	for ; l >= 0 && r <= len(runes)-1 && runes[l] == runes[r]; l, r = l-1, r+1 {
		if r-l+1 >= length { // <= 是因为 OJ 认为 babad 的答案是 aba 而非 bab
			length = r - l + 1
			subStr = string(runes[l : r+1])
		}
	}
	return
}

//
// O(N^3) 复杂度
//
func badLongestPalindrome(s string) string {
	runes := []rune(s)
	n := len(runes)
	if n <= 1 {
		return s
	}

	maxLen := 0
	maxL, maxR := 0, 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPalindrome(runes[i:j+1]) && j-i+1 > maxLen {
				maxLen = j - i + 1
				maxL, maxR = i, j
			}
		}
	}
	return string(runes[maxL : maxR+1]) // 注意对空字符串，转换为 "\u0000"
}

// O(N) 判断是否为回文字符串
func isPalindrome(runes []rune) bool {
	l, r := 0, len(runes)-1
	for l <= r {
		if runes[l] != runes[r] {
			return false
		}
		l++
		r--
	}
	return true
}
