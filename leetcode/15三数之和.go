package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println("res", res)
}

//两数之和 n^3
func threeSum1(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	var left int
	for i := 0; i < len(nums)-2; i++ {
		left = i + 1
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		twoSumList := twoSum1(left, nums)
		if len(twoSumList) > 0 {
			for _, sumList := range twoSumList {
				res = append(res, []int{nums[sumList[0]], nums[sumList[1]], nums[sumList[2]]})

			}
		}
	}
	return res
}

func twoSum1(left int, nums []int) [][]int {
	var res [][]int
	for j := left; j <= len(nums)-2; j++ {
		if j > left && nums[j] == nums[j-1] {
			continue
		}
		for k := j + 1; k <= len(nums)-1; k++ {
			if k > j+1 && nums[k] == nums[k-1] {
				continue
			}
			if nums[left-1]+nums[j]+nums[k] == 0 {
				res = append(res, []int{left - 1, j, k})
			}
		}
		fmt.Println("left", left, "j", j, "res", res)

	}
	return res
}

//两数之和+双指针
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				//避免重复
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}
