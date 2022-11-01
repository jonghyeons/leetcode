func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[last] = i - last
		}
		stack = append(stack, i)
	}
	return res
}