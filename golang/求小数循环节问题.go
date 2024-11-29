package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := DecimalCircle(3, 7)
	fmt.Println("res", res)
}
func DecimalCircle(left int, right int) string {
	var res []int //记录每一次的商
	var mod int
	mod = left % right
	intRes := left / right
	var resStr string
	resStr = strconv.Itoa(intRes)
	resStr += "."
	for mod != 0 {
		res = append(res, mod*10/right)
		mod = mod * 10 % right
		fmt.Println("res", res, "mod", mod)
		if len(res) != 1 && res[len(res)-1] == res[0] {
			break
		}
	}
	if mod != 0 {
		res = res[:len(res)-1]

	}
	var decStr string
	for i := range res {
		decStr += strconv.Itoa(res[i])
	}
	if mod == 0 {
		resStr += decStr
		return resStr
	}
	resStr = resStr + "(" + decStr + ")"
	return resStr
}
