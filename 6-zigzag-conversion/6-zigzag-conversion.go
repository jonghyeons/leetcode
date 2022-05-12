func convert(s string, numRows int) string {
    if numRows == 1 {
		return s
	}
    
	board := make([]string, numRows)
	idx := 0
	topDown := true

	for i := 0; i < len(s); i++ {
		if topDown {
			board[idx] += string(s[i])
			idx++
			if idx == numRows-1 {
				topDown = false
				continue
			}
		} else {
			board[idx] += string(s[i])
			idx--
			if idx <= 0 {
				topDown = true
			}
		}
	}
    
	return strings.Join(board, "")    
}