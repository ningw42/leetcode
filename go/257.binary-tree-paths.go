/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	return addPath("", root)
}

func addPath(prefix string, node *TreeNode) []string {
	if node == nil {
		return []string{}
	} else {
		content := strconv.Itoa(node.Val)
		if prefix != "" {
			content = prefix + "->" + content
		}
		if node.Left != nil && node.Right != nil {
			return append(addPath(content, node.Left), addPath(content, node.Right)...)
		} else if node.Left != nil && node.Right == nil {
			return addPath(content, node.Left)
		} else if node.Left == nil && node.Right != nil {
			return addPath(content, node.Right)
		} else {
			return []string{content}
		}
	}
}

