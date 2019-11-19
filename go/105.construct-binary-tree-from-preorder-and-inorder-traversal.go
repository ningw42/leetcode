/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else {
		var i int
		value := preorder[0]
		for i = 0; i < len(inorder); i++ {
			if inorder[i] == value {
				break
			}
		}
		return &TreeNode{
			Val: value,
			Left: buildTree(preorder[1:i+1], inorder[0:i]),
			Right: buildTree(preorder[i+1:], inorder[i+1:]),
		}
	}
}
// @lc code=end

