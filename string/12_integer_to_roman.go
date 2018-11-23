package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1994)) // MCMXCIV
	fmt.Println(intToRoman(20))   // XX
}

var (
	m = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	units = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

func intToRoman(num int) string {
	var roman string
	for num > 0 {
		for _, unit := range units { // 从大到小一个一个匹配
			if num/unit > 0 {
				roman += m[unit]
				num -= unit
				break // 匹配上一位，进行下一轮 // 20 是 XX 不是 XIXI(10+9+1)
			}
		}
	}
	return roman
}
