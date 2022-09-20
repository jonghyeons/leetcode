func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	var diagonal = make(map[int][]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonal[i-j] = append(diagonal[i-j], mat[i][j])
		}
	}

	for _, v := range diagonal {
		sort.Ints(v)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = diagonal[i-j][0]
			if len(diagonal[i-j]) > 0 {
				diagonal[i-j] = diagonal[i-j][1:]
			}
		}
	}

	return mat
}
