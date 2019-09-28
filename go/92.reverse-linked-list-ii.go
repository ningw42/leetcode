/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	if m == n {
		return head
	}

	var prev, p, next *ListNode
	tail := head
	i := 1
	for ; i < m; i++ {
		prev = tail
		tail = tail.Next
	}

	curr := tail
	for ; i <= n; i++ {
		if p == nil {
			p = curr
			curr = curr.Next
		} else {
			next = curr.Next
			curr.Next = p
			p = curr
			curr = next
		}
	}

	if prev == nil {
		head = p
	} else {
		prev.Next = p
	}
	tail.Next = next

	return head
}

