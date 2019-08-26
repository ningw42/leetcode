/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	ret := 0
	for _, n := range findNumberThroughPath(0, root) {
		ret += n
	}

	return ret
}

func findNumberThroughPath(prev int, root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else {
		value := prev * 10 + root.Val
		if root.Left == nil && root.Right == nil {
			return []int{value}
		} else if root.Left == nil && root.Right != nil {
			return findNumberThroughPath(value, root.Right)
		} else if root.Left != nil && root.Right == nil {
			return findNumberThroughPath(value, root.Left)
		} else {
			return append(findNumberThroughPath(value, root.Left), findNumberThroughPath(value, root.Right)...)
		}
	}
}

