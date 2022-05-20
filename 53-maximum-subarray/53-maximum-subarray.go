func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}

		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}