func plusOne(digits []int) []int {
	var reverse []int
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum := 0

		if i == len(digits)-1 {
			sum = digits[i] + 1
		} else {
			sum = digits[i] + carry
		}

		if sum == 10 {
			sum = 0
			carry = 1
		} else {
			carry = 0
		}
		reverse = append(reverse, sum)
	}
    
    if carry != 0 {
		reverse = append(reverse, 1)
	}
    
	var res []int
	for i := len(reverse) - 1; i >= 0; i-- {
		res = append(res, reverse[i])
	}
	return res
}