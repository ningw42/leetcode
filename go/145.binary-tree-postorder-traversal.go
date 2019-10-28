/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {
	if root != nil {
		return append(
			postorderTraversal(root.Left),
			append(postorderTraversal(root.Right), root.Val)...,
		)
	} else {
		return []int{}
	}
}

// @lc code=end

