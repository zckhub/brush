package main

import "fmt"

func subarraySum(nums []int, k int) int {
	preSumMap := map[int]int{} //前缀和 有多少个前缀和
	preSumMap[0] = 1
	preSum := 0
	res := 0
	for i := range nums {
		//前缀和-k 在preSum出现的次数
		preSum += nums[i]
		res += preSumMap[preSum-k]
		preSumMap[preSum]++

	}
	return res
}

func main() {
	nums := []int{1, 1, 1}
	res := subarraySum(nums, 2)
	fmt.Println(res)
}
