func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 3 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	if len(nums) != 1 && nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}

	return nums[0]
}