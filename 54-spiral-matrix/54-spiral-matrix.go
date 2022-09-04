func spiralOrder(matrix [][]int) []int {
	var res []int
	n, m := len(matrix), len(matrix[0])
	curX, curY := 0, 0
	endX, endY := n-1, m-1

	for curX <= endX && curY <= endY {
		for i := curY; i <= endY; i++ {
			res = append(res, matrix[curX][i])
		}
		curX++

		for i := curX; i <= endX; i++ {
			res = append(res, matrix[i][endY])
		}
		endY--

		if curX <= endX {
			for i := endY; i >= curY; i-- {
				res = append(res, matrix[endX][i])
			}
		}
		endX--

		if curY <= endY {
			for i := endX; i >= curX; i-- {
				res = append(res, matrix[i][curY])
			}
		}
		curY++
	}

	return res
}