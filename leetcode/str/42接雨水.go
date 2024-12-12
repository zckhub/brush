package main

import "fmt"

func main() {
	//height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	height := []int{4, 2, 0, 3, 2, 5}
	res := trap(height)
	fmt.Println("res", res)
}
func trap(height []int) int {
	var slow, fast int
	for i := range height {
		//找到第一个不为零的点
		if height[i] != 0 {
			slow = i
			fast = i + 1
			break
		}
	}
	var res int
	for fast < len(height)-1 {
		minus := 0
		//找到下一个点位 高于当前slow或已到终点
		for fast < len(height)-1 && height[fast] < height[slow] {
			minus += height[fast]
			fast++
		}
		//fast走到终点，slow和fast之间不存在比fast大的
		if fast == len(height)-1 && height[fast] >= height[slow] {
			res += min(height[fast], height[slow])*(fast-slow-1) - minus
			return res
		}
		if fast < len(height)-1 {
			res += min(height[fast], height[slow])*(fast-slow-1) - minus
			//fmt.Println("fast",fast,"slow",slow,"res",res)

			slow = fast
			fast++
		}

	}
	//fast走到终点，slow和fast之间不存在比fast大的。从右往左计算一遍
	right := len(height) - 1
	left := slow
	end := slow

	for j := len(height) - 1; j > slow; j-- {
		if height[j] > 0 {
			right = j
			left = j - 1
			break
		}
	}
	for left > end {
		minus := 0
		for left > end && height[left] < height[right] {
			minus += height[left]
			left--
		}
		res += min(height[left], height[right])*(right-left-1) - minus
		//fmt.Println("left",left,"right",right,"res",res)
		right = left
		left--
	}
	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
