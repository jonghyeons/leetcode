func strStr(haystack string, needle string) int {
    if haystack == needle {
        return 0
    }
    
    length := len(haystack) - len(needle)
    for i := 0; i <= length; i++ {
        if haystack[i:i+len(needle)] == needle {
            return i
        }
    }
    return -1
}