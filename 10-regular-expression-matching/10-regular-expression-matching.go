func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)
	var dp [][]bool

	for i := 0; i <= slen; i++ {
		d := make([]bool, plen+1)
		dp = append(dp, d)
	}

	dp[0][0] = true

	for j := 1; j <= plen; j++ {
		if int(p[j-1]) == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			if s[i-1] == p[j-1] || int(p[j-1]) == '.' {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			} else if int(p[j-1]) == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if s[i-1] == p[j-2] || int(p[j-2]) == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
					dp[i][j] = dp[i][j] || dp[i][j-1]
				}
			}
		}
	}

	return dp[slen][plen]
}