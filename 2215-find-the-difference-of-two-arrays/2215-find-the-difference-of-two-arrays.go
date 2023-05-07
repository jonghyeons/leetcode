func findDifference(nums1 []int, nums2 []int) [][]int {
    m1 := make(map[int]bool)
    m2 := make(map[int]bool)
    for _, num := range nums1 {
        m1[num] = true
    }
    for _, num := range nums2 {
        m2[num] = true
    }
    res := [][]int{[]int{}, []int{}}
    for num, _ := range m1 {
        if _, ok := m2[num]; !ok {
            res[0] = append(res[0], num)
        }
    }
    for num, _ := range m2 {
        if _, ok := m1[num]; !ok {
            res[1] = append(res[1], num)
        }
    }
    return res
}