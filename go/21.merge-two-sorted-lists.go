import "fmt"

/*
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (43.59%)
 * Total Accepted:    427.1K
 * Total Submissions: 979.7K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	head := &ListNode{}
	previous := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			previous.Next = l2
			previous = l2
			l2 = l2.Next
		} else if l1.Val < l2.Val {
			previous.Next = l1
			previous = l1
			l1 = l1.Next
		} else {
			n1 := l1.Next
			n2 := l2.Next

			l1.Next = l2
			previous.Next = l1
			previous = l2

			l1 = n1
			l2 = n2
		}
	}

	if l1 != nil && l2 == nil {
		previous.Next = l1
	} else if l1 == nil && l2 != nil {
		previous.Next = l2
	}

	return head.Next
}

func Print(l *ListNode) {
	for c := l; c != nil; c = c.Next {
		fmt.Println(c.Val)
	}
}
