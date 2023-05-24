package main

import "fmt"

/*
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。



示例 1：

输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。

示例 2:

输入: nums = [1,2,3,4], k = 3
输出: false
*/

func main() {
	nums := []int{3, 3, 10, 2, 6, 5, 10, 6, 8, 3, 2, 1, 6, 10, 7, 2}
	k := 6
	fmt.Println(canPartitionKSubsets(nums, k))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sums := 0
	for _, v := range nums {
		sums += v
	}
	if sums%k != 0 {
		return false
	}
	target := sums / k
	used := make([]bool, len(nums))
	memo := make(map[string]bool) //把used数组转化为字符串，作为key
	var dfs func(int, int, int, []bool) bool
	// start 当前搜索nums内的第start个数　p 当前放置第p个桶，　cur　当前桶内的和
	dfs = func(start int, p int, cur int, used []bool) bool {

		if p == 0 {
			//所有桶放置完毕
			return true
		}
		if cur == target {
			//当前桶放置完毕
			res := dfs(0, p-1, 0, used)
			return res
		}
		for i := start; i < len(nums); i++ {
			if used[i] {
				continue
			}
			//判断i是否可以放入当前桶内
			if cur+nums[i] > target {
				continue
			}
			used[i] = true
			key := toString(used)
			if v, ok := memo[key]; ok {
				if v {
					return true
				} else {
					used[i] = false
					return false
				}
			}
			if dfs(i+1, p, cur+nums[i], used) {
				memo[key] = true
				return true
			}
			memo[key] = false
			used[i] = false
		}
		return false
	}
	return dfs(0, k, 0, used)
}

func toString(used []bool) string {
	s := ""
	for _, v := range used {
		if v {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}
