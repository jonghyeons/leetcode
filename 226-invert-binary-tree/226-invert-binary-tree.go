/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}
    
	var queue []*TreeNode
	queue = append(queue, root)

	for {
		if len(queue) == 0 {
			break
		}
 
        length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

            temp := cur.Left
			cur.Left = cur.Right
			cur.Right = temp
            
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
		}
	}

	return root
}