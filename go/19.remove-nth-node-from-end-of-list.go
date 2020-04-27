/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // create a dummy head node with Next point to head
    // two pointers p1 and p2 point to dummy head
    // p1 moves n steps further
    // then move both p1 and p2
    // when p1 reaches the end
    // p2 is the previous node of the target node
    
    if head == nil {
        return nil
    }
    
    dummy := &ListNode{Next: head}
    var p1, p2 *ListNode
    p1 = dummy
    p2 = dummy
    
    for i := 0; i < n; i++ {
        p1 = p1.Next
    }
    
    for p1.Next != nil {
        p1 = p1.Next
        p2 = p2.Next
    }
   
    p2.Next = p2.Next.Next
    
    return dummy.Next
}

