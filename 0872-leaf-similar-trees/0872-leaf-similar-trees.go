/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leftVal []int
	dfs(root1, &leftVal)
	var rightVal []int
	dfs(root2, &rightVal)

	if len(leftVal) != len(rightVal) {
		return false
	}
	for i := 0; i < len(leftVal); i++ {
		if leftVal[i] != rightVal[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, val *[]int) {
	if isLeaf(root) {
		*val = append(*val, root.Val)
		return
	}
	if root.Left != nil {
		dfs(root.Left, val)
	}

	if root.Right != nil {
		dfs(root.Right, val)
	}
}

func isLeaf(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}
