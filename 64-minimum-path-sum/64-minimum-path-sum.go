func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	sum := make([][]int, m)
	for i := range sum {
		sum[i] = append(sum[i], make([]int, n)...)
	}
	sum[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		sum[0][i] = sum[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if sum[i][j-1]+grid[i][j] < sum[i-1][j]+grid[i][j] {
				sum[i][j] = sum[i][j-1] + grid[i][j]
			} else {
				sum[i][j] = sum[i-1][j] + grid[i][j]
			}
		}
	}

	return sum[m-1][n-1]
}