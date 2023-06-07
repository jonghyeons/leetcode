func minFlips(a int, b int, c int) int {
	var res = 0
	for i := 0; i < 32; i++ {
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1
		bitC := (c >> i) & 1

		if (bitA | bitB) != bitC {
			if bitC == 1 {
				res++
			} else {
				res += bitA + bitB
			}
		}
	}

	return res
}