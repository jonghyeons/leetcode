func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums) - 2
	end := len(nums) - 1
	var closest, distance, d int
	for i := 0; i < length; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, end
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < target {
				d = target - sum
				left++
			} else if sum > target {
				d = sum - target
				right--
			} else {
				return sum
			}
			if d < distance || distance == 0 {
				closest = sum
				distance = d
			}
		}
	}
	return closest
}