func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int)
    for i, v := range nums {
        if n, exist := m[v]; exist && i-n <= k {
            return true
        }
        m[v] = i
    }
    return false
}