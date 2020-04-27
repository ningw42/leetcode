/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    // reverse linked-list
    head = reverse(head)
   
    // add one
    carry := 1
    cursor := head
    prev := &ListNode{Next: head} // dummy head
    for cursor != nil {
        if sum := cursor.Val + carry; sum >= 10 {
            cursor.Val = sum - 10
            carry = 1
        } else {
            cursor.Val = sum
            carry = 0
        }
        cursor = cursor.Next
        prev = prev.Next
    }
    if carry == 1 {
        prev.Next = &ListNode{Val: 1}
    }
    
    // reverse back
    return reverse(head)
}

func reverse(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    curr := head.Next
    prev := head
   
    // single node
    if curr == nil {
        return head
    }
   
    prev.Next = nil
    for curr != nil {
        curr, prev, curr.Next = curr.Next, curr, prev
    }
    
    return prev
}