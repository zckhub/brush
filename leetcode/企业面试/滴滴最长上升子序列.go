package main

import (
	"fmt"
)
func main() {
	nums := []int{1,3,2,4}
	res := newLenght(nums)
	fmt.Println(res)
}

//1,3,2,4 dp[i](第i结尾的最长) = max(dp[j]+1) nums[i]>nums[j]
func maxLength(nums []int)int{
	length := len(nums)
	if length == 0{
		return 0
	}
	dp := make([]int,length)
	for i := range dp{
		dp[i] = 1
	}
	res := 1
	for i := 0;i<length;i++{
		for j := 0;j<i;j++{
			if nums[i] > nums[j]{
				if dp[j]+1 > dp[i]{
					dp[i] = dp[j]+1
				}
			}
		}
		if dp[i] > res{
			res = dp[i]
		}
	}
	return res
}

// 5,1,3,2,4,5,8,2,4,5 tails 最长子序列的最小的结尾
func newLenght(nums []int) int{
	var tails []int
	tails = append(tails, 0)
	for _,num := range nums{
		if len(tails)!=0 && tails[len(tails)-1] < num {
			tails = append(tails, num)
		}else{
			index := searchIndex(tails,num)
			fmt.Println("index",index,"tails",tails)
			tails[index] = num
		}
	}
	return len(tails)-1
}

func searchIndex(nums []int,target int)int{
	l,r := 0,len(nums)
	for l < r{
		mid := (l+r)/2
		if nums[mid] >= target{
			r = mid
		}else{
			l = mid+1
		}
	}
	return l
}