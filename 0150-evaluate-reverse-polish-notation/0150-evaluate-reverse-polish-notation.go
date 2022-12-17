func evalRPN(tokens []string) int {
	operation := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	var stack []int
	for len(tokens) > 0 {
		if _, exist := operation[tokens[0]]; !exist {
			num, _ := strconv.Atoi(tokens[0])
			stack = append(stack, num)
		} else {
			last1 := stack[len(stack)-1]
			last2 := stack[len(stack)-2]
			res := operation[tokens[0]](last2, last1)
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		}
		tokens = tokens[1:]
	}
	return stack[0]
}