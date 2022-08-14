func findTargetSumWays(nums []int, target int) int {
	return dfs(nums, target, 0)
}

func dfs(nums []int, target, sum int) int {
	if len(nums) == 0 {
		if target == sum {
			return 1
		} else {
			return 0
		}
	}
	return dfs(nums[1:], target, sum+nums[0]) + dfs(nums[1:], target, sum-nums[0])
}