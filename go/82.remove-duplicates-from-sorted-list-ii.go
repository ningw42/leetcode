/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	return TwoPass(head)
}

func TwoPass(head *ListNode) *ListNode {
	count := make(map[int]int)
	current := head
	for current != nil {
		if _, exists := count[current.Val]; exists {
			count[current.Val]++
		} else {
			count[current.Val] = 1
		}
		current = current.Next
	}

	var prev *ListNode
	current = head
	for current != nil {
		if c := count[current.Val]; c > 1 {
			if prev == nil {
				head = current.Next
			} else {
				prev.Next = current.Next
			}
			current = current.Next
		} else {
			prev = current
			current = current.Next
		}
	}

	return head
}

