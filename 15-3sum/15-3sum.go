func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
			continue
		}
        
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left += 1
			} else if sum > 0 {
				right -= 1
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
                
                for left < right && nums[left] == nums[left+1] {
					left += 1
				}
                
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
                
				left += 1
				right -= 1
			}
		}
	}
	return res    
}