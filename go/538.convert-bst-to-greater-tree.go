/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
 *
 * https://leetcode.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (49.26%)
 * Total Accepted:    62.9K
 * Total Submissions: 127.5K
 * Testcase Example:  '[5,2,13]'
 *
 * Given a Binary Search Tree (BST), convert it to a Greater Tree such that
 * every key of the original BST is changed to the original key plus sum of all
 * keys greater than the original key in BST.
 *
 *
 * Example:
 *
 * Input: The root of a Binary Search Tree like this:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * Output: The root of a Greater Tree like this:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else {
		convert(root, 0)
		return root
	}
}

func convert(root *TreeNode, s int) {
	if root.Left == nil && root.Right == nil {
		root.Val += s
	} else if root.Left != nil && root.Right == nil {
		convert(root.Left, s+root.Val)
		root.Val += s
	} else if root.Left == nil && root.Right != nil {
		rs := sum(root.Right)
		convert(root.Right, s)
		root.Val += s + rs
	} else {
		rs := sum(root.Right)
		convert(root.Left, s+root.Val+rs)
		convert(root.Right, s)
		root.Val += s + rs
	}
}

func sum(n *TreeNode) int {
	if n == nil {
		return 0
	} else {
		return sum(n.Left) + sum(n.Right) + n.Val
	}
} 