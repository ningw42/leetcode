/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var less, other, lessTail, otherTail *ListNode
	for head != nil {
		if head.Val < x {
			if less == nil {
				less = head
				lessTail = head
			} else {
				lessTail.Next = head
				lessTail = head
			}
		} else {
			if other == nil {
				other = head
				otherTail = head
			} else {
				otherTail.Next = head
				otherTail = head
			}
		}
		head = head.Next
	}

	if otherTail != nil {
		otherTail.Next = nil
	}
	if less == nil {
		return other
	} else {
		lessTail.Next = other
		return less
	}
}

