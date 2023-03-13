/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(left *TreeNode, right *TreeNode) bool {
	if right == left {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
    }

	return isSymmetricTree(left.Left, right.Right) && isSymmetricTree(left.Right, right.Left)
}
