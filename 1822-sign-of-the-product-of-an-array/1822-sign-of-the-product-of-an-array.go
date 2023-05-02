func arraySign(nums []int) int {
    n := 1
    for i := range nums {
        if nums[i] == 0 {
            return 0
        }

        prod := n * nums[i]
        if prod > 0 {
            n = 1
        } else if prod < 0 {
            n = -1
        }
    }
    
    if n > 0 {
        return 1
    } else {
        return -1
    }
}