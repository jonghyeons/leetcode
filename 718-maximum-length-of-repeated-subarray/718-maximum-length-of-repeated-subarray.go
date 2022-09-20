func findLength(nums1 []int, nums2 []int) int {
	max := 0
	var dp = make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		for j := 0; j <= len(nums2); j++ {
			dp[i] = append(dp[i], 0)
		}
	}

	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if dp[i+1][j+1] > max {
					max = dp[i+1][j+1]
				}
			}
		}
	}

	return max
}