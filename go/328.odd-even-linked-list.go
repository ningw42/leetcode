/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    odd := &ListNode{}
    even := &ListNode{}
    oddCurr := odd
    evenCurr := even
    for head != nil {
        // odd
        oddCurr.Next, oddCurr = head, head
        // even
        if head.Next != nil {
            evenCurr.Next, evenCurr, head = head.Next, head.Next, head.Next.Next
        } else {
            head = nil
        }
    }
    evenCurr.Next = nil
    oddCurr.Next = even.Next
    return odd.Next
}