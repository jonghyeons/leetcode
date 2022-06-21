func minimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
        return len(words[i]) > len(words[j])
	})

	res := ""
	for i := range words {
		if i == 0 {
			res = words[0] + "#"
			continue
		}

		if strings.Contains(res, words[i]+"#") {
			continue
		}

		res += words[i] + "#"
	}

	return len(res)
}
