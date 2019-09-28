/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)

	for head != nil {
		if _, e := m[head]; e {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}

