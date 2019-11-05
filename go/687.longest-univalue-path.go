/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
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
var max int

func longestUnivaluePath(root *TreeNode) int {
	max = 0
	getLength(root)
	return max
}

func getLength(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getLength(root.Left)
	right := getLength(root.Right)
	var le, re bool
	if root.Left != nil && root.Left.Val == root.Val {
		left = left + 1
		le = true
	}
	if root.Right != nil && root.Right.Val == root.Val {
		right = right + 1
		re = true
	}

	if left > max {
		max = left
	}
	if right > max {
		max = right
	}
	if le && re && left+right > max {
		max = left + right
	}

	length := 0
	if le && left > length {
		length = left
	}
	if re && right > length {
		length = right
	}

	return length
}

// @lc code=end

