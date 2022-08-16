func firstUniqChar(s string) int {
    res := -1
    m := map[string]int{}
    for i := range s {
        if _, exists := m[string(s[i])]; !exists {
            m[string(s[i])] = 1
        } else {
            m[string(s[i])] = m[string(s[i])] + 1
        }
    }

    for i := range s {
        if m[string(s[i])] == 1 {
            res = i
            break
        }
    }
    return res
}