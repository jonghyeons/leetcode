func uniqueOccurrences(arr []int) bool {
	cnt := map[int]int{}
	for i := range arr {
		if _, exist := cnt[arr[i]]; !exist {
			cnt[arr[i]] = 1
		} else {
			cnt[arr[i]] += 1
		}
	}

	revCnt := map[int]int{}
	for i := range cnt {
		if _, exist := revCnt[cnt[i]]; !exist {
			revCnt[cnt[i]] = i
		} else {
			return false
		}
	}
	return true
}