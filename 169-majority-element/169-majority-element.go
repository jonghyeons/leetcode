func majorityElement(nums []int) int {
    res := 0
    m := map[int]int{}
    
    for _, v := range nums {
        if _, exists := m[v]; exists {
            m[v] += 1
            if m[v] > len(nums)/2 {
                res = v
                break
            }
        } else {
            m[v] = 1
            res = v
        }
    }
    
    return res
}