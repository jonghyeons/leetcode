/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	node := []int{}

	for head != nil {
		node = append(node, head.Val)
		head = head.Next
	}
    
	for len(node) > 1 {
		if node[0] == node[len(node)-1] {
			node = node[1 : len(node)-1]
			continue
		} else {
			return false
		}
	}

    return true
}