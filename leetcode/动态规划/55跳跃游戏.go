package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
func canJump(nums []int) bool {
	var rightAlmost int
	for i := 0; i < len(nums); i++ {
		if i <= rightAlmost {
			rightAlmost = max(rightAlmost, nums[i]+i)
		}
		if rightAlmost >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
