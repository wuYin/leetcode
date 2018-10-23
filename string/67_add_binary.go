package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

// 考虑溢出 ok? 不是 strconv.ParseInt(a, 2, 64) 然后相加就能搞定的...
// 任意进制的任意长度数字相加，需从前向后遍历，求和后取余，并考虑进位
func addBinary(a string, b string) string {
	i1, i2 := len(a)-1, len(b)-1
	res := ""
	carry := 0
	for i1 >= 0 || i2 >= 0 {
		sum := carry
		if i1 >= 0 {
			sum += int(rune(a[i1]) - '0') // '0' == 49 // ASCII
			i1--
		}
		if i2 >= 0 {
			sum += int(rune(b[i2]) - '0')
			i2--
		}

		carry = sum / 2
		sum = sum % 2
		if sum == 0 {
			res += "0"
		} else {
			res += "1"
		}
	}

	if carry > 0 {
		res += "1"
	}

	return reverseStr(res)
}

func reverseStr(s string) string {
	runes := []rune(s)
	n := len(runes)
	mid := n / 2
	for i := 0; i < mid; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
