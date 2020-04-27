/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var slice []*ListNode
    for head != nil {
        slice = append(slice, head)
        head = head.Next
    }
    return slice[len(slice) / 2]
}