func equalPairs(grid [][]int) int {
	var res int
	n := len(grid)
	cols := map[int][]int{}
	for i := 0; i < n; i++ {
		col := []int{}
		for j := 0; j < n; j++ {
			col = append(col, grid[j][i])
		}
		cols[i] = col
	}

	for i := range grid {
		for _, v := range cols {
			if reflect.DeepEqual(grid[i], v) {
				res++
			}
		}
	}

	return res
}