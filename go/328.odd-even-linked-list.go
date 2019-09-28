/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	var odd, even, oddTail, evenTail *ListNode
	count := 1
	for ; head != nil; count++ {
		if count&1 == 1 {
			if odd == nil {
				odd, oddTail = head, head
			} else {
				oddTail.Next, oddTail = head, head
			}
		} else {
			if even == nil {
				even, evenTail = head, head
			} else {
				evenTail.Next, evenTail = head, head
			}
		}
		head = head.Next
	}

	if evenTail != nil {
		evenTail.Next = nil
	}
	if odd == nil {
		return even
	} else {
		oddTail.Next = even
		return odd
	}
}

