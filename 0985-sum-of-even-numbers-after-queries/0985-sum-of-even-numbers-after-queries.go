func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	var res []int
	getEvenSum := func(sl []int) int {
		sum := 0
		for _, v := range sl {
			if v%2 == 0 {
				sum += v
			}
		}
		return sum
	}

	for _, query := range queries {
		nums[query[1]] = nums[query[1]] + query[0]
		res = append(res, getEvenSum(nums))
	}
	return res
}