func maxIceCream(costs []int, coins int) int {
	var res int
	sort.Ints(costs)
	for i := range costs {
		if costs[i] > coins {
			break
		}
		coins -= costs[i]
		res++
	}

	return res
}