/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    cursor := head
    count := 0
    for cursor != nil {
        count++
        cursor = cursor.Next
    }
    k = k % count
    if k == 0 {
        return head
    }
    k = count - k - 1
    cursor = head
    for i := 0; i < k; i++ {
        cursor = cursor.Next
    }
    newHead := cursor.Next
    cursor.Next, cursor = nil, cursor.Next
    for cursor != nil && cursor.Next != nil {
        cursor = cursor.Next
    }
    cursor.Next = head
    return newHead
}