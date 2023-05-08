func diagonalSum(mat [][]int) int {
	var left, right, res int

	left = 0
	right = len(mat[0]) - 1
	for i := range mat {
		if left == right {
			res += mat[i][left]
		} else {
			res += mat[i][left] + mat[i][right]
		}

		left++
		right--
	}

	return res
}