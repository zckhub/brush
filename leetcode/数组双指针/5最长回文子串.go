package main

import "fmt"

func main() {
	res := longestPalindrome("babad")
	fmt.Println(res)
}
func longestPalindrome(s string) string {
	sList := []byte(s)
	if len(sList) == 1 {
		return s
	}
	var res string
	var tmpRes string
	var left, right int
	for i := 0; i < len(sList)-1; i++ {
		left = i
		right = left
		tmpRes = findPalindrome(left, left, s)
		if len(tmpRes) > len(res) {
			res = tmpRes
		}
		right = left + 1
		tmpRes = findPalindrome(left, right, s)
		if len(tmpRes) > len(res) {
			res = tmpRes
		}
	}
	return res
}

func findPalindrome(left int, right int, s string) (res string) {
	sList := []byte(s)
	for left >= 0 && right < len(sList) && sList[left] == sList[right] {
		left--
		right++
	}
	return string(sList[left+1 : right])
}
