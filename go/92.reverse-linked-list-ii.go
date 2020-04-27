/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n {
        return head
    }
    
    var prev, curr, sectionPrev *ListNode
  
    // find section begin
    i := 0
    curr = head
    for i < m {
        curr, prev, sectionPrev = curr.Next, curr, prev
        i++
    }
    
    // reverse linked list
    for curr != nil && i < n {
        curr.Next, curr, prev = prev, curr.Next, curr
        i++
    }
    
    // link section
    if sectionPrev == nil {
        head.Next = curr
        return prev
    } else {
        sectionPrev.Next, sectionPrev.Next.Next = prev, curr
        return head
    }
}