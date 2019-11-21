/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

var Ancestor *TreeNode
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	Ancestor = nil
	containsPQ(root, p, q)
	return Ancestor
}

func containsPQ(root, p, q *TreeNode) (bool, bool) {
	if root == nil {
		return false, false
	} else {
		P := root == p
		Q := root == q
		leftP, leftQ := containsPQ(root.Left, p, q)
		rightP, rightQ := containsPQ(root.Right, p, q)

		P = P || leftP || rightP
		Q = Q || leftQ || rightQ

		if P && Q && Ancestor == nil {
			Ancestor = root
		}

		return P, Q
	}
}
// @lc code=end

