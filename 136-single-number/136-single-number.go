func singleNumber(nums []int) int {
    res := 0
    m := map[int]int{}
    
    for i := range nums {
        if _, exist := m[nums[i]]; !exist {
            m[nums[i]] = 0
        } else {
            delete(m, nums[i])
        }
    }
    
    for i := range m {
       res = i
    }
    
    return res
}