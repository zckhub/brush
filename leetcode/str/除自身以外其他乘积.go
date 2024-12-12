package main

import "fmt"

//[1,2,3,4]  [24,12,8,6]
//nums = [-1,1,0,-3,3] [0,0,9,0,0]
func main() {
	nums := []int{-1, 1, 0, -3, 3}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
func productExceptSelf(nums []int) []int {
	leftNums := make([]int, len(nums))
	leftNums[0] = 1
	rightNums := make([]int, len(nums))
	rightNums[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		leftNums[i] = leftNums[i-1] * nums[i-1]
		rightNums[len(nums)-i-1] = rightNums[len(nums)-i] * nums[len(nums)-i]
	}
	//fmt.Println("leftNums",leftNums,"rightNums",rightNums)
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = leftNums[i] * rightNums[i]
	}
	return ans
}
