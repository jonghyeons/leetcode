func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	sort.Ints(coins)

	for i := 1; i <= amount; i++ {
		max := 10*10*10*10 + 1
		j := 0
		for j < len(coins) && i >= coins[j] {
			if dp[i-coins[j]] < max {
				max = dp[i-coins[j]]
			}
			j++
		}
		dp[i] = max + 1
	}

	if dp[amount] == 10002 {
		return -1
	}
	return dp[amount]
}