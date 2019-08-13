/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := head.Next
	first := head
	second := head.Next
	var last *ListNode
	for ; first != nil && second != nil; {
		var nextFirst, nextSecond *ListNode
		nextFirst = second.Next 
		if nextFirst == nil {
			nextSecond = nil
		} else {
			nextSecond = nextFirst.Next
		}

		if last != nil {
			last.Next = second
		}
		last = first
		second.Next = first

		first = nextFirst
		second = nextSecond
	}
	last.Next = first

	return h
}

