/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    left := inorderTraversal(root.Left)
    right := inorderTraversal(root.Right)

    res := []int{}
    
    res = append(res, left...)
    res = append(res, root.Val)
    res = append(res, right...)
    
    return res    
}