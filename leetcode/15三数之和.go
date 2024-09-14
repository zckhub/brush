package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4, 0, 1}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := ThreeSum(nums)
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
func ThreeSum(nums []int) (res [][]int) {
	for start := 0; start < len(nums)-2; start++ {
		if start >= 1 && nums[start] == nums[start-1] {
			continue
		}
		fmt.Println("start", start, "-nums[start]", -nums[start])
		resTwoSum := TwoSum(-nums[start], start+1, nums)
		if len(resTwoSum) > 0 {
			res = append(res, resTwoSum...)
		}
	}
	return
}

func TwoSum(target int, start int, nums []int) (res [][]int) {
	right := len(nums) - 1
	for left := start; left < len(nums)-1; left++ {
		if left >= 1 && nums[left] == nums[left-1] {
			continue
		}
		for right > left && nums[left]+nums[right] > target {
			right--
		}

		if nums[left]+nums[right] == target {
			res = append(res, []int{-target, nums[left], nums[right]})
		}

	}
	return
}

//两数之和+双指针
func threeSum2(nums []int) [][]int {
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
