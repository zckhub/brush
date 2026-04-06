package main

import "fmt"

func main(){
	nums := []int{0,0,1,1,1,1,2,3,3}
	res := removeDuplicates(nums)
	fmt.Println(res)
	fmt.Println(nums)
}
func removeDuplicates(nums []int)int{
	if len(nums)<=2{
		return len(nums)
	}
	slow :=2
	fast :=2
	ans := 2
	for fast <len(nums){
		if nums[slow-2] != nums[fast]{
			nums[slow] = nums[fast]
			slow++
			ans++
		}
		fast++
	}
	return ans
}
func removeDuplicates2(nums []int) int {
	left := 1
	record := 1
	ans := 1
    for i := 1;i<len(nums);i++{
		if nums[i] == nums[i-1]{
			if record <2{
				nums[left] = nums[i]
				left++
				ans++
			}
			record++
		}else{
			nums[left] = nums[i]
			left++
			record = 1
			ans++
		}
	}
	return ans
}