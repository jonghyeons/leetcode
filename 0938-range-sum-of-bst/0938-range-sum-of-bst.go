/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBST(root *TreeNode, low int, high int) int {
	var sum int
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	if root.Left != nil {
		sum += rangeSumBST(root.Left, low, high)
	}

	if root.Right != nil {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}