func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
    
	memo := make(map[string][]string)
    
	var backtrack func(s string) []string
	backtrack = func(s string) []string {
		if result, found := memo[s]; found {
			return result
		}
        
		var result []string
        
		if s == "" {
			return []string{""}
		}
        
		for i := 1; i <= len(s); i++ {
			word := s[:i]
			if wordSet[word] {
				remaining := s[i:]
				subResults := backtrack(remaining)
				for _, subResult := range subResults {
					if subResult == "" {
						result = append(result, word)
					} else {
						result = append(result, word+" "+subResult)
					}
				}
			}
		}
        
		memo[s] = result
		return result
	}
    
	return backtrack(s)
}