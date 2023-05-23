package main

import "fmt"

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：

输入：nums = [1]
输出：[[1]]
*/
func main() {

	nums := []int{1, 2, 3}
	res := nsolution(nums)
	fmt.Println("res = ", res)

}

func nsolution(nums []int) [][]int {
	n := len(nums)
	used := make([]bool, n)
	track := make([]int, 0)
	var res [][]int
	var backtrack func()
	backtrack = func() {
		if len(track) == n {
			fmt.Println("track = ", track)
			tmp := make([]int, n)
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			track = append(track, nums[i])
			backtrack()
			used[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return res
}
