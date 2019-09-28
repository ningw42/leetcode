import "sort"

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	i := 1
	for ; i < len(nodes); i++ {
		nodes[i-1].Next = nodes[i]
	}
	nodes[i-1].Next = nil

	return nodes[0]
}

