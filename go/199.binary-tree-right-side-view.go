/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 *
 * https://leetcode.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (45.51%)
 * Total Accepted:    140.7K
 * Total Submissions: 308K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * Given a binary tree, imagine yourself standing on the right side of it,
 * return the values of the nodes you can see ordered from top to bottom.
 *
 * Example:
 *
 *
 * Input: [1,2,3,null,5,null,4]
 * Output: [1, 3, 4]
 * Explanation:
 *
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else {
		if root.Left == nil && root.Right == nil {
			return []int{root.Val}
		} else if root.Left != nil && root.Right == nil {
			return append([]int{root.Val}, rightSideView(root.Left)...)
		} else if root.Left == nil && root.Right != nil {
			return append([]int{root.Val}, rightSideView(root.Right)...)
		} else {
			l := rightSideView(root.Left)
			r := rightSideView(root.Right)
			if len(l) > len(r) {
				return append(append([]int{root.Val}, r...), l[len(r):]...)
			} else {
				return append([]int{root.Val}, r...)
			}
		}
	}
}