func search(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for {
		if start > end {
			break
		}

		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[start] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < nums[end] {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if nums[start] == nums[mid] {
				start++
			}
			if nums[end] == nums[mid] {
				end--
			}
		}
	}
	return -1
}