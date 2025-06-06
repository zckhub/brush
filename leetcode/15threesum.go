package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4, 0, 1}
	nums := []int{-2, 0, 1, 1, 2}
	//-4,-1,-1,0,0,1,1,2
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	res := GeneralThreeSum(0, nums)
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
		//fmt.Println("left", left, "j", j, "res", res)

	}
	return res
}
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	for start := 0; start < len(nums)-2; start++ {
		if start >= 1 && nums[start] == nums[start-1] {
			continue
		}
		//fmt.Println("start", start, "-nums[start]", -nums[start])
		resTwoSum := TwoSum(-nums[start], start+1, nums)
		if len(resTwoSum) > 0 {
			res = append(res, resTwoSum...)
		}
	}
	return res
}

func TwoSum(target int, start int, nums []int) (res [][]int) {
	right := len(nums) - 1
	for left := start; left < len(nums)-1; left++ {
		if left > start && nums[left] == nums[left-1] {
			continue
		}
		for right > left && nums[left]+nums[right] > target {
			right--
		}

		if right > left && nums[left]+nums[right] == target {
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

func GeneralTwoSum(target int, start int, nums []int) [][]int {
	left := start
	right := len(nums) - 1
	var res [][]int
	for left < right {
		if nums[left]+nums[right] == target {
			res = append(res, []int{nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left-1] == nums[left] {
				left++
			}
			for left < right && nums[right+1] == nums[right] {
				right--
			}
		}
		for left < right && nums[left]+nums[right] < target {
			left++
		}
		for left < right && nums[left]+nums[right] > target {
			right--
		}
		//fmt.Println("left",left,"right",right,"t",target)
	}
	return res
}

func GeneralThreeSum(target int, nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		twoResList := GeneralTwoSum(target-nums[i], i+1, nums)
		for index := range twoResList {
			res = append(res, append(twoResList[index], nums[i]))
		}
	}
	return res
}

/*
2024.12.27
func threeSum(nums []int) [][]int {
	sort.Slice(nums,func(i,j int)bool{
		return nums[i]<nums[j]
	})
	var res [][]int
	for i := range nums{
		if i!=0 && nums[i] == nums[i-1]{
			continue
		}
		tmpRes := twoSum(nums,i)
		if len(tmpRes) != 0{
			res = append(res,tmpRes...)
		}
	}
	return res
}

//nums[i]和后面的两个数之和为0
func twoSum(nums []int, start int)[][]int{
	if start >= len(nums)-2{
		return [][]int{}
	}
	var res [][]int
	target := -nums[start]
	left,right:=start+1,len(nums)-1
	for left<right{
		for left>start+1&&left<right && nums[left]==nums[left-1]{
			left++
		}
		if left!=right && nums[left]+nums[right]==target{
			//这里已经实现了去重，所以该函数不用right去重
			//fmt.Println("left",left,"right",right,"start",start)
			res = append(res,[]int{nums[start], nums[left],nums[right]})
			left++
			right--
		}
		if nums[left]+nums[right]>target{
			right--
		}
		if nums[left]+nums[right]<target{
			left++
		}
	}
	return res

}
*/
