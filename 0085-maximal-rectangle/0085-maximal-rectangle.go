func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows := len(matrix)
	cols := len(matrix[0])
	maxArea := 0
	heights := make([]int, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		var curHeight int
		if i == len(heights) {
			curHeight = 0
		} else {
			curHeight = heights[i]
		}

		for len(stack) > 0 && curHeight < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			maxArea = max(maxArea, area)
		}
		stack = append(stack, i)
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}