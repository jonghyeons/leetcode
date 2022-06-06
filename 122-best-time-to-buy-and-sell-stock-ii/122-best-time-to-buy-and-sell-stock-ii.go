func maxProfit(prices []int) int {
	profit := 0
	for i := range prices {
		if i != len(prices)-1 && prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}

	}
	return profit    
}