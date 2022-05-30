func removeElement(nums []int, val int) int {
	length := len(nums)
	pivot := 0

	for length > pivot {
		if nums[pivot] == val {
			nums = append(nums[0:pivot], nums[pivot+1:]...)
			length--
			continue
		}
		pivot++
	}

	return length
}