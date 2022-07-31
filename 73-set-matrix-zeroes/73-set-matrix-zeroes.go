func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	var geo [][]int
	for i := range matrix {
		for j, v := range matrix[i] {
			if v == 0 {
				geo = append(geo, []int{i, j})
			}
		}
	}

	for _, v := range geo {
		x := v[0]
		y := v[1]
		for i := x; i >= 0; i-- {
			matrix[i][y] = 0
		}
		for i := x; i < m; i++ {
			matrix[i][y] = 0
		}

		for i := y; i >= 0; i-- {
			matrix[x][i] = 0
		}
		for i := y; i < n; i++ {
			matrix[x][i] = 0
		}
	}
}