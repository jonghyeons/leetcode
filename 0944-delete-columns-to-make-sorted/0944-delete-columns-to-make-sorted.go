func minDeletionSize(strs []string) int {
	var res int
	for i := 0; i < len(strs[0]); i++ {
		prev := strs[0][i]
		for j := 1; j < len(strs); j++ {
			next := strs[j][i]
			if prev > next {
				res++
				break
			} else {
				prev = next
			}
		}
	}
	return res
}