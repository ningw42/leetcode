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
	if n == 0 {
		return head
	}
	
	stash := make([]*ListNode, 0)

	current := head
	for current != nil {
		stash = append(stash, current)
		current = current.Next
	}

	length := len(stash)
	if length == 0 {
		return nil
	} else if length == 1 {
		return nil
	} else {
		if n == 1 {
			stash[length - 2].Next = nil
			return head
		} else if n >= length {
			return head.Next
		} else {
			stash[length - n - 1].Next = stash[length - n + 1]
			return head
		}
	}
}

