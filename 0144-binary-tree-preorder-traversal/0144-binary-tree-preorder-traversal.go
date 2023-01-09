/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		left := preorderTraversal(root.Left)
		right := preorderTraversal(root.Right)
		res = append(res, root.Val)
		res = append(res, left...)
		res = append(res, right...)
	}

	return res
}