package main

import "fmt"

// gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// gas = [2,3,4], cost = [3,4,3]
//输入：nums = [0,1,2,2,3,0,4,2], val = 2
//[3,2,2,3], val = 3
//输出：5, nums = [0,1,4,0,3,_,_,_]
func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res := removeElement(nums, val)
	fmt.Println("res", res, "nums", nums)
}
func removeElement1(nums []int, val int) int {
	var left, right int
	right = len(nums) - 1
	if right == 0 {
		if nums[right] == val {
			return right
		}
	}
	for left < right {
		for left < right && nums[right] == val {
			right--
		}
		for left < right && nums[left] != val {
			left++
		}
		if left < right {
			nums[left] = nums[right]
			right--
			left++
		}

	}
	if right >= 0 && nums[right] != val {
		return right + 1
	}
	return right
}

func removeElement(nums []int, val int) int {
	var slow int
	for fast := range nums {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
