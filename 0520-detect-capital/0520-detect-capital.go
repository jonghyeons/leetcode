func detectCapitalUse(word string) bool {
	if word[0] >= 65 && word[0] <= 90 {
		if len(word) == 1 {
			return true
		}

		if word[0] >= 65 && word[1] <= 90 {
			for i := 1; i < len(word); i++ {
				if word[i] >= 97 && word[i] <= 122 {
					return false
				}
			}
		} else {
			for i := 1; i < len(word); i++ {
				if word[i] >= 65 && word[i] <= 90 {
					return false
				}
			}
		}
	} else {
		for i := 1; i < len(word); i++ {
			if word[i] >= 65 && word[i] <= 90 {
				return false
			}
		}
	}
	return true
}