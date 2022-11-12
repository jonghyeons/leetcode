func makeGood(s string) string {
	if s == "" {
		return s
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1]-32 || s[i] == s[i+1]+32 {
			return makeGood(s[:i] + s[i+2:])
		}
	}
	return s    
}