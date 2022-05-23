func lengthOfLastWord(s string) int {
    arr := strings.Split(s, " ")
    idx := len(arr)-1
    for {
        if arr[idx] != "" {
            return len(arr[idx])
        } else {
            idx--
        }
    }
}