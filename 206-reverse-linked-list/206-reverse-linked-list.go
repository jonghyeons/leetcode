/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    return reverse(head, nil)
}

func reverse(node, prev *ListNode) *ListNode{
    if node == nil {
		return prev
	}
    
	next := node.Next
	node.Next = prev
    
    return reverse(next, node)
}