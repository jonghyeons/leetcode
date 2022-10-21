func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for i := range words {
		if _, exist := m[words[i]]; exist {
			m[words[i]] = m[words[i]] + 1
		} else {
			m[words[i]] = 1
		}
	}

	rm := map[int][]string{}
	for key, v := range m {
		if _, exist := rm[v]; !exist {
			rm[v] = []string{key}
		} else {
			rm[v] = append(rm[v], key)
		}
	}

	var keys []int
	for key := range rm {
		keys = append(keys, key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var res []string
	for i := range keys {
		if _, exist := rm[keys[i]]; exist {
			sort.Strings(rm[keys[i]])
			res = append(res, rm[keys[i]]...)
		}
	}

	return res[:k]
}