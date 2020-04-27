/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var current, previous *ListNode
    current = head.Next
    previous = head
    for current != nil {
        current.Next, current, previous = previous, current.Next, current
    }
    head.Next = nil
    return previous
}