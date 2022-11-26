func subArrayRanges(nums []int) int64 {
    var sum int64
    var max int   
    for i, v := range nums {
        max = v
        for _, n := range nums[i+1:] {
            if max < n {
                max = n
            }
            if v > n {
                v = n
            }
            sum += int64(max-v)
        }
    }
    return sum 
}