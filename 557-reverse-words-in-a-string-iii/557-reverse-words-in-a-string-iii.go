func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i, word := range words {
		sl := strings.Split(word, "")
		for i := 0; i < len(sl)/2; i++ {
			sl[i], sl[len(sl)-1-i] = sl[len(sl)-1-i], sl[i]
		}
		words[i] = strings.Join(sl, "")
	}

	return strings.Join(words, " ")
}