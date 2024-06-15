func minIncrementForUnique(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    
    sort.Ints(nums)
    moves := 0
    prev := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] <= prev {
            moves += prev - nums[i] + 1
            prev++
        } else {
            prev = nums[i]
        }
    }
    
    return moves
}