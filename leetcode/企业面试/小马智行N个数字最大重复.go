package main

import (
	"fmt"
	"sort"
)
// 给 N 个数字和一个 K， 允许最多执行 K 次操作，每次他选择任何一个数字加一。问如何操作时的最终重复数字最多？输出这个数字以及最大重复次数。

func main() {
	num, cnt := maxRepeat([]int{1, 2, 3, 4}, 3)
	fmt.Println("数字:", num, " 最大重复次数:", cnt)
}

func maxRepeat(nums []int, k int) (int, int) {
	sort.Ints(nums)
	n := len(nums)
	maxCnt := 1
	resNum := nums[0]
	add := 0
	left := 0

	for right := 0; right < n; right++ {
		if right > 0 {
			add += (nums[right] - nums[right-1]) * (right - left)
		}

		// 操作次数超了，缩小窗口
		for add > k {
			add -= nums[right] - nums[left]
			left++
		}

		// 更新最大重复
		cnt := right - left + 1
		if cnt > maxCnt {
			maxCnt = cnt
			resNum = nums[right]
		}
	}
	return resNum, maxCnt
}


