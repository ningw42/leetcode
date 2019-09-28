/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	return buildTree(nodes)
}

func buildTree(nodes []*ListNode) *TreeNode {
	length := len(nodes)
	if length == 0 {
		return nil
	}
	mid := length / 2
	return &TreeNode{
		Val:   nodes[mid].Val,
		Left:  buildTree(nodes[0:mid]),
		Right: buildTree(nodes[mid+1:]),
	}
}

