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
