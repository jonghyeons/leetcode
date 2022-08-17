func uniqueMorseRepresentations(words []string) int {
	trans := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morse := map[string]bool{}

	for _, word := range words {
		transformation := ""
		for i := range word {
			transformation += trans[word[i]-97]
		}
		morse[transformation] = true
	}
	return len(morse)
}