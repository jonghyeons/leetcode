func groupAnagrams(strs []string) [][]string {
	var res [][]string
	sorting := func(str string) string {
		arr := strings.Split(str, "")
		sort.Strings(arr)
		return strings.Join(arr, "")
	}

	m := map[string]int{}
	for i := range strs {
		sorted := sorting(strs[i])
		if _, exist := m[sorted]; !exist {
			res = append(res, []string{strs[i]})
			m[sorted] = len(res) - 1
		} else {
			res[m[sorted]] = append(res[m[sorted]], strs[i])
		}
	}
	return res
}
