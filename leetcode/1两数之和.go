package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 11, 15, 7}
	sum := 9
	res := twoSum(nums, sum)
	fmt.Println(res)
}

func twoSum(nums []int, sum int) []int {
	helpMap := map[int]int{}
	var res []int
	for i := range nums {
		if p, ok := helpMap[sum-nums[i]]; ok {
			res = append(res, i)
			res = append(res, p)
			return res
		}
		helpMap[nums[i]] = i
	}
	return res
}
