func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := 0; i <= k+1; i++ {
		dp[src][i] = 0
	}

	for i := 1; i <= k+1; i++ {
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			if dp[from][i-1] != math.MaxInt32 {
				dp[to][i] = min(dp[to][i], dp[from][i-1]+price)
			}
		}
	}

	result := math.MaxInt32
	for i := 1; i <= k+1; i++ {
		result = min(result, dp[dst][i])
	}

	if result == math.MaxInt32 {
		return -1
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}