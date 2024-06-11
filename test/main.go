package main

import (
	"fmt"
	"strconv"
	"strings"
)

//输入年月日，输出是当年的第几天
func main() {
	inputStr := "2024-03-01"
	res := getDays(inputStr)
	fmt.Println("res", res)
}

func getDays(inputStr string) int {
	var res int
	strList := strings.Split(inputStr, "-")
	if len(strList) < 3 {
		return res
	}
	towMouthDay := 29
	year, _ := strconv.Atoi(strList[0])
	month, _ := strconv.Atoi(strList[1])
	day, _ := strconv.Atoi(strList[2])
	if year%4 != 0 {
		towMouthDay = 28
	}
	monthDays := map[int]int{
		1:  31,
		2:  towMouthDay + 31,
		3:  towMouthDay + 31 + 31,
		4:  towMouthDay + 31 + 31 + 30,
		5:  towMouthDay + 31 + 31 + 30 + 31,
		6:  towMouthDay + 31 + 31 + 30 + 31 + 30,
		7:  towMouthDay + 31 + 31 + 30 + 31 + 30 + 31,
		8:  towMouthDay + 31 + 31 + 30 + 31 + 30 + 31 + 31,
		9:  towMouthDay + 31 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
		10: towMouthDay + 31 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
		11: towMouthDay + 31 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
		12: towMouthDay + 31 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
	}
	res = monthDays[month-1] + day

	return res
}
