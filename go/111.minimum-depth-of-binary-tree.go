/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
		return 0
	} else {
		if root.Left == nil && root.Right == nil {
			return 1
		} else if root.Left == nil && root.Right != nil {
			return minDepth(root.Right) + 1
		} else if root.Left != nil && root.Right == nil {
			return minDepth(root.Left) + 1
		} else {
			l := minDepth(root.Left)
			r := minDepth(root.Right)
			if l > r {
				return r + 1
			} else {
				return l + 1
			}
		}
	}
}

