var dp map[int]int

func climbStairs(n int) int {
    if dp == nil {
        dp = make(map[int]int)    
    }
    
	if n <= 2 {
		return n
	}

	if _, exist := dp[n]; exist {
		return dp[n]
	}

	dp[n] = climbStairs(n-1) + climbStairs(n-2)
	return dp[n]
}