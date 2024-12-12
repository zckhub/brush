package main

import "fmt"

//[7,3,5,1,12,2,11]
func main() {
	prices := []int{7, 3, 5, 1, 12, 2, 11}
	res := maxProfit(prices)
	fmt.Println("res", res)
}
func maxProfit(prices []int) int {
	var res int
	left, right := 0, len(prices)-1
	for left < right {
		for left < right && prices[left+1] <= prices[left] {
			left++
		}
		for left < right && prices[right-1] >= prices[right] {
			right--
		}
		res = max(res, prices[right]-prices[left])
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a

}
