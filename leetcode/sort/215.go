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

func FindKth(nums []int, k int) int {
	return findKthHelp(nums, k, 0, len(nums)-1)
}
func findKthHelp(nums []int, k int, start int, end int) int {
	index := partion1(nums, start, end)
	if index == k-1 {
		return nums[index]
	}
	if index > k-1 {
		return findKthHelp(nums, k, start, index-1)
	}
	return findKthHelp(nums, k-index-1, index+1, end)
}
func partion1(nums []int, start int, end int) int {
	value := nums[start]
	left := start + 1
	right := end
	if left > right {
		return start
	}
	for left <= right {
		for left < right && nums[left] >= value {
			left++
		}
		for left <= right && nums[right] <= value {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[start] = nums[start], nums[left]
	return left
}
