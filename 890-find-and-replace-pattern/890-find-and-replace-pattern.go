func findAndReplacePattern(words []string, pattern string) []string {
	var res []string
	numPattern := getNumericPattern(pattern)

	for _, word := range words {
		wordPattern := getNumericPattern(word)
		if numPattern == wordPattern {
			res = append(res, word)
		}
	}
	return res
}

func getNumericPattern(word string) string {
	m := map[string]int{}
	idx := 0
	pattern := ""
	for _, s := range word {
		if _, exists := m[string(s)]; !exists {
			idx++
			m[string(s)] = idx
			pattern += strconv.Itoa(idx)
		} else {
			pattern += strconv.Itoa(m[string(s)])
		}
		pattern += "/"
	}
	return pattern
}