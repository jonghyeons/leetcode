/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	level := 1
	maxSum := math.MinInt32
	max := 0

	for len(queue) > 0 {
		size := len(queue)
		levelSum := 0

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
            
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if levelSum > maxSum {
			maxSum = levelSum
			max = level
		}

		level++
	}

	return max
}