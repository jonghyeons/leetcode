func topKFrequent(nums []int, k int) []int {
	type frequency struct {
		num  int
		freq int
	}

	var res []int
	m := map[int]int{}

	for _, v := range nums {
		if _, exist := m[v]; !exist {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}

	var f []frequency
	for i := range m {
		f = append(f, frequency{i, m[i]})
	}

	sort.Slice(f, func(i, j int) bool {
		return f[i].freq > f[j].freq
	})

	for i := 0; i < k; i++ {
		res = append(res, f[i].num)
	}

	return res
}