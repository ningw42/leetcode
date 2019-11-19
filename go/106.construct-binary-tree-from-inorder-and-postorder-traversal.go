/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	} else {
		var i int
		value := postorder[len(postorder)-1]
		for i = 0; i < len(inorder); i++ {
			if inorder[i] == value {
				break
			}
		}
		return &TreeNode{
			Val: value,
			Left: buildTree(inorder[0:i], postorder[0:i]),
			Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
		}
	}
}
// @lc code=end

