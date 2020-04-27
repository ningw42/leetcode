/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	var current, prev, next *ListNode
	dummy := &ListNode{Next:nil}
	unsorted := head
	for unsorted != nil {
		prev, current = dummy, dummy.Next
		for current != nil && current.Val < unsorted.Val {
			prev = current
			current = current.Next
		}
		prev.Next = unsorted
		next = unsorted.Next
		unsorted.Next = current
		unsorted = next
	}
	return dummy.Next
}