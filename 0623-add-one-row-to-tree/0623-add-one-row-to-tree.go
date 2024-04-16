/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val, depth int) *TreeNode {
    if depth == 1 {
        newNode := &TreeNode{Val: val, Left: root}
        return newNode
    }
    addRow(root, val, depth-1)
    return root
}

func addRow(node *TreeNode, val, depth int) {
    if node == nil {
        return
    }
    if depth == 1 {
        left := node.Left
        right := node.Right
        node.Left = &TreeNode{Val: val, Left: left}
        node.Right = &TreeNode{Val: val, Right: right}
    } else {
        addRow(node.Left, val, depth-1)
        addRow(node.Right, val, depth-1)
    }
}