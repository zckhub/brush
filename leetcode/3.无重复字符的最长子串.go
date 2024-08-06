package main

import "fmt"

func main() {
	res := lengthOfLongestSubstring("dvdf")
	fmt.Println(res)
}
func lengthOfLongestSubstring(s string) int {
	help := map[byte]int{}
	left, right, res, tmpRes := 0, 0, 0, 0
	for right <= len(s)-1 {
		if right <= len(s)-1 && help[s[right]] == 0 {
			help[s[right]]++
			right++
			tmpRes++
			if res <= tmpRes {
				res = tmpRes
			}
		} else {
			help[s[left]]--
			left++
			tmpRes--
		}

	}
	return res
}
