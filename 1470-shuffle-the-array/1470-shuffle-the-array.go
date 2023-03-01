func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		if i%2 == 0 {
			res[i] = nums[i/2]
		} else {
			res[i] = nums[n]
			n++
		}
	}
	return res
}