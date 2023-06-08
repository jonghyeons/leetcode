func countNegatives(grid [][]int) int {
	var res int
	for _, row := range grid {
		left, right := 0, len(row)-1
		for left <= right {
			if row[left] > -1 {
				left++
			} else if row[right] > -1 {
				right--
			} else {
				res += right - left + 1
				break
			}
		}
	}
	return res
}