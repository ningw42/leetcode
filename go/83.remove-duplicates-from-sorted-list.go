/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	prev := head
	current := head.Next
	for current != nil {
		for current != nil && current.Val == prev.Val {
			current = current.Next
		}
		prev.Next = current
		prev = current
		if current != nil {
			current = current.Next
		}
	}

	return head
}

