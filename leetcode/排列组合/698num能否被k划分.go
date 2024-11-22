package main

import "fmt"

func main() {
	nums1 := []int{4, 3, 2, 3, 5, 2, 1}
	res1 := canPartitionKSubsets(nums1, 4)
	fmt.Println(res1)
	nums2 := []int{1, 2, 3, 4}
	res2 := canPartitionKSubsets(nums2, 3)
	fmt.Println(res2)
}

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	var sum int
	for _, value := range nums {
		sum += value
	}
	//k个桶的集合
	buckets := make([]int, k)

	if sum%k != 0 {
		return false
	}

	target := int(sum / k)
	return backtrack(nums, 0, buckets, target)
}

func backtrack(nums []int, index int, buckets []int, target int) bool {
	if index == len(nums) {
		//fmt.Println(target,nums,buckets)
		//检查所有桶数字是否都为target
		for _, value := range buckets {
			if value != target {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(buckets); i++ {
		//给index所在位置的数字找到桶
		//桶装满了
		if buckets[i]+nums[index] > target {
			continue
		}

		buckets[i] += nums[index]
		res := backtrack(nums, index+1, buckets, target)
		//所有桶已经装满，返回成功
		if res == true {
			return true
		}

		buckets[i] -= nums[index]

	}
	//nums[index]装入哪个桶都不行
	return false
}
