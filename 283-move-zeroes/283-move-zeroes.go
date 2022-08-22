func moveZeroes(nums []int) {
    if len(nums) == 1 {
        return
    }

    left, right := 0, 0
    
    for left = 0; left < len(nums); left++ {
        if nums[left] != 0 {
            if left != right {
                nums[left], nums[right] = nums[right], nums[left]
            }
            right++
        }
    }
}