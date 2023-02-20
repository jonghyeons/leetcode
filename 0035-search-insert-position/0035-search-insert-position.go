func searchInsert(nums []int, target int) int {
    start, end := 0, len(nums)-1
    
    for start <= end {
        mid := (start + end + 1)/2 
        if nums[mid] < target {
            start = mid + 1
        } else if nums[mid] > target {
            end = mid - 1
        } else {
            return mid
        }
    }
    
    return start
}