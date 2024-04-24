package main

import (
	"fmt"
)

func NumberTostring(number int) string {
	var resStr string
	for number > 0 {
		resStr = string(number%1000) + resStr
		number = number / 1000
	}
	return resStr
}

func main() {
	var input int
	fmt.Print("请输入整数: ")
	fmt.Scanln(&input)
	num1 := input/10 + 119105013026
	str1 := NumberTostring(num1)
	num2 := num1 - 118983996991
	str2 := NumberTostring(num2)
	num3 := num1 + 108982009006013
	str3 := NumberTostring(num3)
	num4 := input - 841725
	str4 := NumberTostring(num4)
	fmt.Println(str1 + " " + str2 + " " + str3 + " " + str4)
	//waitSignal()
}
