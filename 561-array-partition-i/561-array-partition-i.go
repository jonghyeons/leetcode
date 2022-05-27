func arrayPairSum(nums []int) int {
    res := 0
    sort.Ints(nums)
    for i := range nums {
        if i%2 != 0 {
            res += nums[i-1]
        }
    }
    return res
}