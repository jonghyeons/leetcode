func jump(nums []int) int {
	res, max, maxJump := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		if maxJump < i+nums[i] {
			maxJump = i + nums[i]
		}

		if i == max {
			max = maxJump
			res++
		}
	}
	return res
}
