func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	hash := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	num, last, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = hash[char]
		if num < last {
			res -= num
		} else {
			res += num
		}
		last = num
	}

	return res
}