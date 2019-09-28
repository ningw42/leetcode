/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	return Origin(head)
}

func FollowUp(head *ListNode) bool {
	// TODO
	return false
}

func Origin(head *ListNode) bool {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	for i, j := 0, len(list)-1; i < j; {
		if list[i] != list[j] {
			return false
		}
		i++
		j--
	}

	return true
}

