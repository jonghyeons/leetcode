/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	var list []int

	tmp := head
	for {
		if tmp == nil {
			break
		}
		list = append(list, tmp.Val)
		tmp = tmp.Next
	}

	sort.Ints(list)

	tmp = head
	for i := range list {
		tmp.Val = list[i]
		tmp = tmp.Next
	}
    
	return head
}