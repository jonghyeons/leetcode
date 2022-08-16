func isPalindrome(s string) bool {
    ss := ""
    for i := range s {
        if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
            ss += strings.ToLower(string(s[i]))
        }
    }

    for i := 0; i < len(ss); i++ {
        if ss[i] != ss[len(ss)-1-i] {
            return false
        }        
    }
    return true
}