func addBinary(a string, b string) string {
	if len(a) > len(b) {
		gap := len(a) - len(b)
		for i := 0; i < gap; i++ {
			b = "0" + b
		}
	} else if len(b) > len(a) {
		gap := len(b) - len(a)
		for i := 0; i < gap; i++ {
			a = "0" + a
		}
	}
    
	arrA := strings.Split(a, "")
	arrB := strings.Split(b, "")

	carry := 0
    var arr []string
	for i := len(arrA) - 1; i >= 0; i-- {
		numA, _ := strconv.Atoi(arrA[i])
		numB, _ := strconv.Atoi(arrB[i])

		sum := numA + numB + carry
		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
		arr = append(arr, strconv.Itoa(sum%2))
	}
    
	if carry != 0 {
		arr = append(arr, strconv.Itoa(1))
	}

	var reverse []string
	for i := len(arr) - 1; i >= 0; i-- {
		reverse = append(reverse, arr[i])
	}

	return strings.Join(reverse, "")    
}