package main

import "fmt"

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	res := 0
	var visitedRight byte
	visited := map[byte]int{}
	sList := []byte(s)
	for right < len(sList) {
		visitedRight = sList[right]
		visited[sList[right]]++
		right++

		for visited[visitedRight] > 1 {
			visited[sList[left]]--
			left++
		}
		if res < right-left {
			res = right - left
		}
		//fmt.Println("left",left,"right",right)
	}
	return res
}
