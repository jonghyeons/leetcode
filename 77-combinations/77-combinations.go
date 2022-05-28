var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	dfs([]int{}, n, k, 1)
	return res
}

func dfs(prev []int, n, k, start int) {
	if 0 == k {
		res = append(res, prev)
		return
	}

    elements := make([]int, len(prev))
    copy(elements, prev)
    
	for i := start; i <= n; i++ {
        dfs(append(elements, i), n, k-1, i+1)
	}
}
