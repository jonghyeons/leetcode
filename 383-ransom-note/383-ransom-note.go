func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := map[string]int{}

	for _, v := range magazine {
		if _, exists := magazineMap[string(v)]; !exists {
			magazineMap[string(v)] = 1
		} else {
			magazineMap[string(v)]++
		}
	}

	for _, v := range ransomNote {
		if _, exists := magazineMap[string(v)]; !exists {
			return false
		} else {
			if magazineMap[string(v)] > 1 {
				magazineMap[string(v)]--
			} else {
				delete(magazineMap, string(v))
			}
		}
	}

	return true
}