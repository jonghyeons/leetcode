func reverse(x int) int {
	res := ""
	str := strconv.Itoa(x)

	if string(str[0]) == "-" {
		res = "-"
		str = str[1:]
	} else if string(str[len(str)-1]) == "0" {
		str = str[:len(str)-1]
	}

	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}

	n, _ := strconv.Atoi(res)

	if n < -1<<31 || n > 1<<31-1 {
		return 0
	}
	
	return n
}