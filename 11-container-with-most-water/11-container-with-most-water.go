func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		w := right - left
		h := 0
		if height[left] <= height[right] {
			h = height[left]
		} else {
			h = height[right]
		}

		if max < w*h {
			max = w * h
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
