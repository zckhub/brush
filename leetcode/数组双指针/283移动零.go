package main

func moveZeroes(nums []int) {
	var cnt int
	for _, num := range nums {
		if num != 0 {
			nums[cnt] = num
			cnt++
		}
	}
	for i := cnt; i <= len(nums)-1; i++ {
		nums[i] = 0
	}

}
