func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		if left > right {
			return -1
		}
		mid := (left + right) / 2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
}