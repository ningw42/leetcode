/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	m := make(map[int][]int)
	Traversal(root, 0, m)

	ret := make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		ret[len(m) - i - 1] = m[i]
	}

	return ret
}

func Traversal(root *TreeNode, level int, m map[int][]int) {
	if root != nil {
		Traversal(root.Left, level+1, m)
		Traversal(root.Right, level+1, m)
		m[level] = append(m[level], root.Val)
	}
}
// @lc code=end

