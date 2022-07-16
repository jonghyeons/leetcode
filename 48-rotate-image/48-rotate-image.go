func rotate(matrix [][]int) {
	max := len(matrix) - 1
	res := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		res[i] = append(res[i], matrix[i]...)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][max-i] = res[i][j]
		}
	}
}
