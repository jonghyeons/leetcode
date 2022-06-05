func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		nums[len(nums)-1], nums[len(nums)-2] = nums[len(nums)-2], nums[len(nums)-1]
		return
	}

	startIdx := len(nums) - 1
	for i := startIdx; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
		startIdx = i - 1
	}

	if startIdx == 0 {
		sort.Sort(sort.IntSlice(nums))
		return
	}

	min := nums[startIdx]
	swapIdx := startIdx
	for i := startIdx + 1; i < len(nums); i++ {
		if nums[i] < min && nums[i] > nums[startIdx-1] {
			min = nums[i]
			swapIdx = i
		}
	}

	nums[swapIdx], nums[startIdx-1] = nums[startIdx-1], nums[swapIdx]
	sort.Sort(sort.IntSlice(nums[startIdx:]))
}