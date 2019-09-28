/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	var prev *ListNode
	for node.Next != nil {
		node.Val = node.Next.Val
		prev = node
		node = node.Next
	}

	prev.Next = nil
}

