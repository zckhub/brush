package main

import "fmt"

//贪心从左往右遍历，发现大的话 +1 保证了右边一定不小于左边的，
//再从右往左，保证左边的也不会小于右边
func candy(ratings []int) int {
	resList := make([]int, len(ratings))
	resList[0] = 1
	var res int
	for i := 1; i <= len(ratings)-1; i++ {
		resList[i] = 1
		if ratings[i] > ratings[i-1] {
			resList[i] = resList[i-1] + 1
		}
	}
	for j := len(ratings) - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] && resList[j] <= resList[j+1] {
			resList[j] = resList[j+1] + 1
		}
	}
	for k := range resList {
		res += resList[k]
	}
	return res
}

func main() {
	ra := []int{1, 3, 4, 5, 2}
	fmt.Println(candy(ra))
}
