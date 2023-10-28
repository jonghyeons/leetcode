func longestPalindrome(s string) string {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		l1 := twoPointer(i, i, s)
		l2 := twoPointer(i, i+1, s)
		length := l1
		if l2 > l1 {
			length = l2
		}
		if length > right-left {
			left = i - (length-1)/2
			right = i + length/2
		}
	}
	return s[left : right+1]
}

func twoPointer(left, right int, s string) int {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return right - left - 1
}