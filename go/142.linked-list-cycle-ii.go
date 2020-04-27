/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    return twoPointers(head)
}

func twoPointers(head *ListNode) *ListNode {
    i := findIntersection(head)
    if i == nil {
        return nil
    }
    j := head
    for i != j {
        i = i.Next
        j = j.Next
    }
    return i
}

func findIntersection(head *ListNode) *ListNode {
    i := head
    j := head
    for i != nil && j != nil {
        i = i.Next
        j = j.Next
        if j != nil {
            j = j.Next
        }
        if i == j {
            return i
        }
    }
    return nil
}