func mergeAlternately(word1 string, word2 string) string {
    var res string
    len1, len2 := len(word1), len(word2)
    for i := 0; i < len1 || i < len2; i++ {
        if i < len1 {
            res += string(word1[i])
        }
        if i < len2 {
            res += string(word2[i])
        }
    }
    return res
}