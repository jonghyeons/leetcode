func generateParenthesis(n int) []string {
	var res []string
	dfs(&res, []string{}, n, n)
	return res
}

func dfs(res *[]string, prev []string, left, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, strings.Join(prev, ""))
		return
	}

	if left > 0 {
		dfs(res, append(prev, "("), left-1, right)
	}

	if right > 0 && left < right {
		dfs(res, append(prev, ")"), left, right-1)
	}
}
