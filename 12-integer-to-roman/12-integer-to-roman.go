func intToRoman(num int) string {
	res := ""
	idx := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	hash := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	
	for num > 0 {
		for i := range idx {
			if idx[i] <= num {
				res += hash[idx[i]]
				num -= idx[i]
				break
			} else {
				continue
			}
		}
	}

	return res
}