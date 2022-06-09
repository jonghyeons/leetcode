func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[0]) - 1

	for col >= 0 && row <= len(matrix) - 1 {
		if matrix[row][col] < target {
			row++
		} else if matrix[row][col] > target {
			col--
		} else if matrix[row][col] == target {
			return true
		}
	}

	return false    
}