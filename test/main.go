package main

import (
	"fmt"
	"sort"
)

func main() {
	// var a int
	// fmt.Scan(&a)
	nums := []int{-5, 1, -3, -1, -4, -2, 4, -1, -1, 4}
	fmt.Println(threeSum(nums))
}
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	for i := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		tmpRes := twoSum(nums, i)
		if len(tmpRes) != 0 {
			res = append(res, tmpRes...)
		}
	}
	return res
}

//nums[i]和后面的两个数之和为0
func twoSum(nums []int, start int) [][]int {
	if start >= len(nums)-2 {
		return [][]int{}
	}
	var res [][]int
	target := -nums[start]
	left, right := start+1, len(nums)-1
	for left < right {

		if nums[left]+nums[right] == target {
			//这里已经实现了去重，所以该函数不用right去重
			//fmt.Println("left",left,"right",right,"start",start)
			res = append(res, []int{nums[start], nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				//去重原则：判断同样的情况前一个已经出现了，直接++
				left++
			}
			//for left<right && nums[right] == nums[right+1]{
			//	right--
			//}
		}
		if nums[left]+nums[right] > target {
			right--
		}
		if nums[left]+nums[right] < target {
			left++
		}
	}
	return res

}
