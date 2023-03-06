func findKthPositive(arr []int, k int) int {
    var idx int
    idx = sort.Search(len(arr), func(i int) bool{
        missing := arr[i]-i-1 
        return missing >= k 
    })
    
    return idx + k
}