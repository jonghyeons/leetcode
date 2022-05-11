func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i, v := range nums {
        m[v] = i
    }
    
    for i, v := range nums {
        if j, ok := m[target-v]; ok && i != j {
            return []int{i, j}
        }
    }
    
    return []int{}
}