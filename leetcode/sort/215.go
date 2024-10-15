package sort

func findKthLargest(nums []int, k int) int {
	return help(nums, 0, len(nums)-1, k-1)
}

func help(nums []int, left int, right int, k int) int {

	index := partion(nums, left, right)
	if index == k {
		return nums[index]
	}
	if index < k {
		return help(nums, index+1, right, k)
	}
	if index > k {
		return help(nums, left, index-1, k)
	}
	return nums[index]
}

//从大到小排序

func partion(nums []int, start int, end int) int {

	value := nums[start]
	for start < end {
		for start < end && nums[end] <= value {
			end--
		}
		nums[start] = nums[end]

		for start < end && nums[start] >= value {
			start++
		}
		nums[end] = nums[start]
	}
	//此时start = end
	nums[start] = value
	//fmt.Println("start = ", start, "end = ", end, "nums = ", nums)
	return start

}
