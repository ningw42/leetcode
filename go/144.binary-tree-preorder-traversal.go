/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root != nil {
		return append(
			[]int{root.Val},
			append(preorderTraversal(root.Left), preorderTraversal(root.Right)...)...,
		)
	} else {
		return []int{}
	}
}

// @lc code=end

