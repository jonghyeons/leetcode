/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	swap := true
	node := head

	for node.Next != nil {
		if swap {
			node.Val, node.Next.Val = node.Next.Val, node.Val
		}
		swap = !swap
		node = node.Next
		continue
	}

	return head
}
