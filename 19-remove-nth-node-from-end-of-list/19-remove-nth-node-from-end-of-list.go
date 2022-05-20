/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        return nil
    }
    
    tmp := &ListNode{0, head}
	left := tmp
	right := tmp

	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right != nil && right.Next != nil {
		right = right.Next
		left = left.Next
	}
	
    if left.Next != nil {
        left.Next = left.Next.Next    
    } 
    
	return tmp.Next
}