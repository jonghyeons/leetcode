func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    visited := make([]bool, len(nums))
    permutation := []int{}
    
    sort.Ints(nums)

    backtrack(&res, nums, &permutation, &visited)  
    return res
}

func backtrack(res *[][]int, nums []int, permutation *[]int, visited *[]bool) {
    if len(*permutation) == len(nums) {
        temp := make([]int, len(*permutation))
        copy(temp, *permutation)
        *res = append(*res, temp)
        return
    }
    
    for i := 0; i < len(nums); i++ {
        if (*visited)[i] || (i > 0 && nums[i] == nums[i-1] && !(*visited)[i-1]) {
            continue
        }
        
        *permutation = append(*permutation, nums[i])
        (*visited)[i] = true
        
        backtrack(res, nums, permutation, visited)
        
        *permutation = (*permutation)[:len(*permutation)-1]
        (*visited)[i] = false
    }
}
