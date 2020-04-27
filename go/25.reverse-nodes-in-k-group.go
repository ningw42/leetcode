/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    var last, cursor, current, prev *ListNode
    var count int
    
    cursor = head
    for cursor != nil {
        count = 0
        // make sure there are more than k nodes left
        for cursor != nil && count < k {
            cursor = cursor.Next
            count++
        }
        if count == k {
            // preparation for linked list traversal
            if last == nil {
                // the very first iteration
                prev = head
                current = head.Next
            } else {
                prev = last.Next
                current = last.Next.Next
            }
            // reverse linked list for current section
            for current != cursor {
                current, prev, current.Next = current.Next, current, prev
            }
            // put the reversed section in right place
            if last == nil {
                // the very first iteration
                head, head.Next, last = prev, current, head
            } else {
                last.Next, last, last.Next.Next = prev, last.Next, current
            }
        }
    }
    
    return head
}