func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    res := append(nums1, nums2...)
    
    sort.Ints(res)
    
    if len(res)%2 != 0 {
        return float64(res[len(res)/2])
    } else {
        return float64(res[len(res)/2-1] + res[len(res)/2])/2
    }
}