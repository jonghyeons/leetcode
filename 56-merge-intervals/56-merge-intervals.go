func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		prev := res[len(res)-1]
		if prev[1] >= intervals[i][0] {
			if prev[1] < intervals[i][1] {
				prev[1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}