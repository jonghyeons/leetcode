func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	check := map[interface{}]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	res := false
	stack := list.New()
	for i := range s {
		if _, exist := check[string(s[i])]; exist {
			stack.PushFront(string(s[i]))
			continue
		} else {
			if stack.Len() == 0 {
				return false
			}
			if check[stack.Front().Value] == string(s[i]) {
				stack.Remove(stack.Front())
				res = true
				continue
			} else {
				return false
			}
		}
	}

	if stack.Len() != 0 {
		return false
	} else {
		return res
	}
}