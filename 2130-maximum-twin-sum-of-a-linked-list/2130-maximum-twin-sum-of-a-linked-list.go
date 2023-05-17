/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	var res int
	node := head
	cnt := 1
	for {
		if node.Next != nil {
			node = node.Next
			cnt++
		} else {
			break
		}
	}

	mid := cnt / 2
	node = head
	for i := 0; i < mid; i++ {
		node = node.Next
	}

	reverse := reverseNode(node, nil)
	for reverse != nil {
		sum := reverse.Val + head.Val
		if sum > res {
			res = sum
		}
		reverse = reverse.Next
		head = head.Next
	}
	return res
}

func reverseNode(node, prev *ListNode) *ListNode {
	if node == nil {
		return prev
	}

	next := node.Next
	node.Next = prev

	return reverseNode(next, node)
}