func mostCommonWord(paragraph string, banned []string) string {
	reg := regexp.MustCompile(`[!?''.,;]`)
	s := reg.ReplaceAllString(paragraph, " ")
	arr := strings.Split(s, " ")
	m := map[string]int{}
	for i := range arr {
		word := strings.ToLower(arr[i])
		if word == "" {
			continue
		}
		if _, exists := m[word]; !exists {
			m[word] = 1
		} else {
			m[word]++
		}
	}

	contains := func(s string, arr []string) bool {
		for i := range arr {
			if arr[i] == s {
				return true
			}
		}
		return false
	}

	type Count struct {
		Key string
		Cnt int
	}

	var c []Count
	for k, v := range m {
		if contains(k, banned) {
			delete(m, k)
		} else {
			c = append(c, Count{k, v})
		}
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i].Cnt > c[j].Cnt
	})

	return c[0].Key
}