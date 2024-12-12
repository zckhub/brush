package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	strg := "-91283472332"
	res := myAtoi(strg)
	fmt.Println("res", res)
}

func myAtoi(s string) int {
	var res int
	var sign string
	for i := range s {
		if sign != "" && (s[i] < '0' || s[i] > '9') {
			break
		}
		if s[i] == byte(' ') {
			continue
		}
		if sign == "" && s[i] == '+' {
			sign = "+"
		}
		if sign == "" && s[i] == '-' {
			sign = "-"
		}
		if (s[i] != '-' && s[i] != '+') && (s[i] < '0' || s[i] > '9') {
			break
		}
		if s[i] >= '0' && s[i] <= '9' {
			if sign == "" {
				sign = "+"
			}
			tmp, _ := strconv.Atoi(string(s[i]))
			res = res*10 + tmp
			if sign == "+" && res > math.MaxInt32 {
				return math.MaxInt32
			}
			if sign == "-" && res > math.MaxInt32 {
				return math.MinInt32
			}
			fmt.Println("res", res, "i", tmp)
		}

	}
	if sign == "-" {
		return -res
	}
	return res
}
