func addToArrayForm(num []int, k int) []int {
	var res []int
	var s1, s2 string
	for i := range num {
		s1 += strconv.Itoa(num[i])
	}
	s2 = strconv.Itoa(k)

	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			s2 = "0" + s2
		}
	} else {
		for i := 0; i < len2-len1; i++ {
			s1 = "0" + s1
		}
	}

	carry := 0
	for i := len(s1) - 1; i >= 0; i-- {
		n1, _ := strconv.Atoi(string(s1[i]))
		n2, _ := strconv.Atoi(string(s2[i]))

		res = append([]int{(n1 + n2 + carry) % 10}, res...)
		if n1+n2+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}

	}

	if carry == 1 {
		res = append([]int{1}, res...)
	}

	return res
}