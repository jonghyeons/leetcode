func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
	max := make([]int, len(nums))
	max[0] = nums[0]
	if nums[0] > nums[1] {
		max[1] = nums[0]
	} else {
		max[1] = nums[1]
	}

	for i := 2; i < len(nums); i++ {
		if max[i-1] > nums[i]+max[i-2] {
			max[i] = max[i-1]
		} else {
			max[i] = nums[i] + max[i-2]
		}
	}

	return max[len(max)-1]
}