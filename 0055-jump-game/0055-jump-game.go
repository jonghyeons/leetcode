func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	max := 0
	for i := range nums {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}