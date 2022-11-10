func reverseVowels(s string) string {
	vowels := map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
		"A": 0,
		"E": 0,
		"I": 0,
		"O": 0,
		"U": 0,
	}
	var res string
	var words []string
	for i := range s {
		if _, exist := vowels[string(s[i])]; exist {
			words = append(words, string(s[i]))
		}
	}

	for i := range s {
		if _, exist := vowels[string(s[i])]; exist {
			res += words[len(words)-1]
			words = words[:len(words)-1]
		} else {
			res += string(s[i])
		}
	}

	return res
}