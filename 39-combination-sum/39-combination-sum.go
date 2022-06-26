func combinationSum(candidates []int, target int) [][]int {
    res := &[][]int{}
    sort.Ints(candidates)
    dfs([]int{}, candidates, target, res)
    return *res
}

func dfs(cur, candidates []int, target int, res *[][]int) {
    if target == 0 {
        *res = append(*res, cur)
    }
    
    if len(candidates) == 0 || target < candidates[0]{
        return
    }
    
    cur = cur[:len(cur):len(cur)]
    
    dfs(append(cur, candidates[0]), candidates, target-candidates[0], res)
    dfs(cur, candidates[1:], target, res)
}
