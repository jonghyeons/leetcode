func searchMatrix(matrix [][]int, target int) bool {
	x := 0
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == target {
			return true
		} else if matrix[i][0] > target {
			break
		} else if matrix[i][0] < target {
			x = i
		}
	}

	for i := 0; i < len(matrix[x]); i++ {
		if matrix[x][i] == target {
			return true
		}
	}

	return false
}