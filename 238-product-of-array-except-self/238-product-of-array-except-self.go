func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	tmp := 1
	for i := range nums {
		res[i] = tmp
		tmp *= nums[i]
	}

	tmp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= tmp
		tmp *= nums[i]
	}

	return res
}