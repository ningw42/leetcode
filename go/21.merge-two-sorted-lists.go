import "fmt"

/*
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (43.59%)
 * Total Accepted:    427.1K
 * Total Submissions: 979.7K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
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
