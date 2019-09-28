/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return Reverse(l1, l2)
}

func NonReverse(l1, l2 *ListNode) *ListNode {
	// TODO
	return nil
}

func Reverse(l1, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	var prev *ListNode
	head := l1
	carry := 0
	for {
		if l1 == nil && l2 == nil {
			break
		} else if l1 != nil && l2 == nil {
			sum := l1.Val + carry
			l1.Val = sum % 10
			carry = sum / 10
			prev = l1
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			sum := l2.Val + carry
			l2.Val = sum % 10
			carry = sum / 10
			prev.Next = l2
			prev = l2
			l2 = l2.Next
		} else {
			sum := l1.Val + l2.Val + carry
			l1.Val = sum % 10
			carry = sum / 10
			prev = l1
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	if carry > 0 {
		prev.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return reverseList(head)
}

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

