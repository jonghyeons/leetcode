func kthSmallest(matrix [][]int, k int) int {
	var sum []int
	for i := range matrix {
		sum = append(sum, matrix[i]...)
	}
	sort.Ints(sum)
	return sum[k-1]    
}