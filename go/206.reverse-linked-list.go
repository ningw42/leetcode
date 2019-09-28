/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var dummy *ListNode

	for head != nil {
		next := head.Next
		head.Next = dummy
		dummy = head
		head = next
	}

	return dummy
}

