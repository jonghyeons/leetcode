func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	lengths := make([]int, n)
	counts := make([]int, n)

	for i := 0; i < n; i++ {
		lengths[i] = 1
		counts[i] = 1
	}

	maxLen := 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if lengths[j]+1 > lengths[i] {
					lengths[i] = lengths[j] + 1
					counts[i] = counts[j]
				} else if lengths[j]+1 == lengths[i] {
					counts[i] += counts[j]
				}
			}
		}

		if lengths[i] > maxLen {
			maxLen = lengths[i]
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		if lengths[i] == maxLen {
			result += counts[i]
		}
	}

	return result
}