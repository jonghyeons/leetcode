func lengthOfLongestSubstring(s string) int {
	res, left := 0, 0
	h := map[string]int{}
	for i, v := range s {
		if _, exist := h[string(v)]; exist && left <= h[string(v)]  {
			left = h[string(v)] + 1
		} else {
			if i-left+1 > res {
				res = i - left + 1
			}
		}
		h[string(v)] = i
	}
	return res
}