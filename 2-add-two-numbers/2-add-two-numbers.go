/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
	tmp := res
	carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		nextVal := (val1 + val2 + carry) % 10
		carry = (val1 + val2 + carry) / 10

		tmp.Next = &ListNode{Val: nextVal, Next: nil}
		tmp = tmp.Next
	}
    
	return res.Next
}
