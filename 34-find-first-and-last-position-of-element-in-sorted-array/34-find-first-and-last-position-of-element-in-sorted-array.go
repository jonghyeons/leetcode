func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	mid := len(nums) / 2

	from, to := -1, -1
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			from, to = mid, mid
			break
		}
	}

	if from == -1 && to == -1 {
		return []int{from, to}
	}

	for i := from - 1; i >= 0; i-- {
		if nums[i] == target {
			from = i
		} else if nums[i] < target {
			break
		}
	}

	for i := to + 1; i < len(nums); i++ {
		if nums[i] == target {
			to = i
		} else if nums[i] > target {
			break
		}
	}
	return []int{from, to}
}