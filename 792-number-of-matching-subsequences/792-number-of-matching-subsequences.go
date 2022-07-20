func numMatchingSubseq(s string, words []string) int {
    res := 0
    for i := range words {
        if isSubSeq(s, words[i]) {
            res++
        }
    }
    
    return res
}

func isSubSeq(s, word string) bool{
    idx := 0
    for i := 0; i < len(s); i++ {
        if idx == len(word) {
            return true
        }
        
        if s[i] == word[idx] {
            idx++
        }
    }
    return idx == len(word)
}