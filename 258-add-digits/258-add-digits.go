func addDigits(num int) int {
	var digit int
	for {
		s := strconv.Itoa(num)
		if len(s) == 1 {
			digit, _ = strconv.Atoi(s)
			break
		}
		sum := 0
		for i := range s {
			n, _ := strconv.Atoi(string(s[i]))
			sum += n
		}
		num = sum
	}
	return digit
}