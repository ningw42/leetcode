/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    for l1 != nil || l2 != nil {
        if l1 != nil && l2 != nil {
            if l1.Val < l2.Val {
                current.Next, current, l1 = l1, l1, l1.Next
            } else {
                current.Next, current, l2 = l2, l2, l2.Next
            }
        } else if l1 == nil && l2 != nil {
            current.Next, current, l2 = l2, l2, l2.Next
        } else if l1 != nil && l2 == nil {
            current.Next, current, l1 = l1, l1, l1.Next
        }
    }
    
    return dummy.Next
}