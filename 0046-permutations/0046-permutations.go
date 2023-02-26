var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	dfs([]int{}, nums)
	return res
}

func dfs(prev, nums []int) {
	if len(nums) == 0 {
		res = append(res, prev)
		return
	}

	for i, v := range nums {
		element := make([]int, len(nums))
		copy(element, nums)
		dfs(append(prev, v), append(element[:i], element[i+1:]...))
	}
}