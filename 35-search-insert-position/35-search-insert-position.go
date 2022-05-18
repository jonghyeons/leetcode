func searchInsert(nums []int, target int) int {
    for i := range nums {
        if nums[i] == target {
            return i
        }
    }
    
    nums = append(nums, target)
    sort.Ints(nums)
    
    idx := 0
    for i := range nums {
        if nums[i] == target {
            idx = i
            break
        }
    }
    return idx
}