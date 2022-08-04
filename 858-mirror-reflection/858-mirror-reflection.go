func mirrorReflection(p int, q int) int {
	i := lcm(p, q)
	if (i/q)%2 == 0 {
		return 2
	}

	return (i / p) % 2
}

func lcm(p, q int) int {
	return p * q / gcd(p, q)
}


func gcd(p, q int) int {
	for q != 0 {
		p, q = q, p%q
	}

	return p
}