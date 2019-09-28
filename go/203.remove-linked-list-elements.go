/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		if curr.Val == val {
			if prev == nil {
				head = curr.Next
				curr = curr.Next
			} else {
				prev.Next = curr.Next
				curr = curr.Next
			}
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return head
}

