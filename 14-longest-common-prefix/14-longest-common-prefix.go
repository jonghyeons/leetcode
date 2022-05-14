func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
	}
    
	res := ""
	idx := 0
	for {
		if idx > 200 {
			break
		}
		for i, _ := range strs {
            if len(strs[i]) < idx {
				return res
			}
            
			if !strings.HasPrefix(strs[i], strs[0][0:idx]) {
				return res
			}
		}
		res = strs[0][0:idx]
		idx++
	}

	return res    
}