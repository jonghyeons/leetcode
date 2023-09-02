var res []string

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
    
    res = []string{}
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
    
	dfs(m, "", digits, 0)
	return res
}

func dfs(m map[string]string, word, digits string, index int) {
	if index == len(digits) {
		res = append(res, word)
		return
	}
    
	for i := range m[string(digits[index])] {
		dfs(m, word+string(m[string(digits[index])][i]), digits, index+1)
	}
}