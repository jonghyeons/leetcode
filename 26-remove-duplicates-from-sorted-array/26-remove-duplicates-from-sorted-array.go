func removeDuplicates(nums []int) int {
	length := len(nums)
    idx := 1
    
	for i := 1; i < length; i++ {
		if nums[i-1] != nums[i] {
			nums[idx] = nums[i]
			idx++
		}
	}
    
	return idx
}