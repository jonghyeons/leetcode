func minSetSize(arr []int) int {
    key := []int{}
    m := make(map[int]int)
    length := len(arr)
    
    for _, v := range arr {
       m[v]++
    }
    
    for i := range m {
        key = append(key, m[i])
    }
    
    sort.Slice(key, func(i,j int) bool {
        return key[i] > key[j]
    })
    
    for i, v := range key {
        length -= v
        if length <= len(arr)/2 {
            return i + 1
        } 
    }
    return len(arr)
}