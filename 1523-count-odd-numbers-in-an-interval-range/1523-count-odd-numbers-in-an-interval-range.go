func countOdds(low int, high int) int {
	var res int
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			res++
		}
	}
	return res
}