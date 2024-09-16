package main

/*
res = max(min(i,j)*s)
每次i或j 移动时，由于s必然要-1，结果值是以小的为准，如果让大的移动，那么最大只能是小的，结果值必然变小
所以要选则min(i,j)移动，才有可能让结果值变大
*/
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	for left < right {
		if height[left] <= height[right] {
			res = max(res, (right-left)*height[left])
			left++
		}
		if height[left] > height[right] {
			res = max(res, (right-left)*height[right])
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
