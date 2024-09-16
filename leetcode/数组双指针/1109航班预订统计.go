package main

import "fmt"

func main() {
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	n := 5
	res := corpFlightBookings(bookings, n)
	fmt.Println(res)
}

//本题需要记知识点
//原理：差分数组的前缀和是原数组
//差分数组和原数组一一对应,本题根据差分数组求原数组
//当数组多个连续元素同时加减时，用差分数组
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, value := range bookings {
		left := value[0] - 1
		right := value[1]
		seats := value[2]
		diff[left] += seats
		if right < n {
			diff[right] -= seats
		}
	}
	nums := make([]int, n)
	nums[0] = diff[0]
	for i := 1; i < n; i++ {
		nums[i] = diff[i] + nums[i-1]
	}
	return nums
}
