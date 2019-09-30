/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}

	var nodes []*ListNode
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}

	index := len(nodes) - 1
	i, j := 0, index
	for i < j {
		if j < index {
			nodes[j+1].Next = nodes[i]
		}
		nodes[i].Next = nodes[j]
		i++
		j--
	}
	if i == j {
		nodes[j+1].Next = nodes[i]
	}
	nodes[i].Next = nil
}

