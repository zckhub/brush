package main

import (
	"fmt"
	"math"
)

func main() {
	// res := findMidNum()
	fmt.Println("hello")
}

//[2,5,6,7,8,34,76]  [3,5,7,8,43]
//2,5,6 | 7  8 34 76
//3,5,7 |  8,43

//1,2  3,4

func findMidNum(nums1 []int, nums2 []int) float64 {
	low := 0           //nums1左边最少放0个
	high := len(nums1) // nums1左边最多放high个.决定左边到底放几个
	m := len(nums1)
	n := len(nums2)
	//左边要放的总数
	leftCnt := (m + n + 1) / 2
	for low <= high {
		// nums1 左边放i个
		i := (low + high) / 2
		//nums2 左边放j个
		j := leftCnt - i
		tmpRightMin := math.MaxInt //右边最小
		tmpLeftMax := math.MinInt  //左边最大
		tmpRightMinB := math.MaxInt
		tmpLeftMaxB := math.MinInt
		if i < m {
			tmpRightMinB = nums1[i]
		}
		if j < n {
			tmpLeftMaxB = nums2[j]
		}
		//满足切割条件-左边都<=右边
		if tmpLeftMax <= tmpRightMinB && tmpLeftMaxB <= tmpRightMin {
			if (m+n)%2 == 1 {
				return max(tmpLeftMax, tmpLeftMaxB)
			}
			leftMax := max(tmpLeftMax, tmpLeftMaxB)
			rightMin := min(tmpRightMinB, tmpRightMin)
			return (leftMax + rightMin) / 2

		} else if tmpLeftMax > tmpRightMinB {
			high = i - 1
		} else {
			low = i + 1
		}

	}
	return 0.0
}
func max(a, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}

func min(a, b int) float64 {
	if a > b {
		return float64(b)
	}
	return float64(a)
}
