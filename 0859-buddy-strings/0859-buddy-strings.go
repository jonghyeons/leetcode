func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	hash := make(map[byte]int)
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			count++
			if count > 2 {
				return false
			}
		}

		hash[s[i]]++
		hash[goal[i]]--
	}

	if s == goal && len(hash) == len(s) {
		return false
	}

	for i := range hash {
		if hash[i] != 0 {
			return false
		}
	}

	return true
}