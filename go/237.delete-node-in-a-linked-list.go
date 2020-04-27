/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    // since node will not be the tail
    // we can copy the next node's value to node
    // and delete the next node
    // O(1)
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}