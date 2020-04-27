/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else {
		return append(inorderTraversal(root.Left), append([]int{root.Val}, inorderTraversal(root.Right)...)...)
	}
}
