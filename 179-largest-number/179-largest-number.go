func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
		str1 := strconv.Itoa(nums[i])
		str2 := strconv.Itoa(nums[j])
		if str1+str2 > str2+str1 {
			return true
		} else {
			return false
		}
	})

	if nums[0] == 0 {
		return "0"
	}

	res := ""
	for i := range nums {
		res += strconv.Itoa(nums[i])
	}
    
	return res
}