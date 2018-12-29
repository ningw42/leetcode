/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 *
 * https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (39.93%)
 * Total Accepted:    207K
 * Total Submissions: 516.6K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * Given a binary tree, flatten it to a linked list in-place.
 *
 * For example, given the following tree:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 *
 *
 * The flattened tree should look like:
 *
 *
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
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
func flatten(root *TreeNode) {
	do(root)
}

func do(r *TreeNode) *TreeNode {
	if r == nil {
		return nil
	} else {
		if r.Left == nil && r.Right == nil {
			return r
		} else if r.Left == nil && r.Right != nil {
			return do(r.Right)
		} else if r.Left != nil && r.Right == nil {
			r.Right = r.Left
			r.Left = nil
			return do(r.Right)
		} else {
			rt := r.Right
			r.Right = r.Left
			r.Left = nil
			do(r.Right).Right = rt
			return do(rt)
		}
	}
}